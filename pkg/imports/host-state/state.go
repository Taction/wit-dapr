package host_state

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	contribMetadata "github.com/dapr/components-contrib/metadata"
	comstate "github.com/dapr/components-contrib/state"
	"github.com/taction/wit-dapr/dapr/state"
	"github.com/taction/wit-dapr/pkg/common"
	"github.com/taction/wit-dapr/pkg/imports/host-state/registery"
	"github.com/tetratelabs/wazero/api"
	"sync"
)

const GetLen = 48
const SetLen = 64

type HostImpl struct {
	lock              sync.RWMutex
	componentSpec     map[string]Component
	componentInstance map[string]comstate.Store
}

var State = &HostImpl{}

func AddComp(comps ...Component) {
	for _, comp := range comps {
		State.componentSpec[comp.Name] = comp
	}
}

func (h *HostImpl) getComponent(ctx context.Context, name string) (comstate.Store, bool) {
	h.lock.Lock()
	defer h.lock.Unlock()
	c, ok := h.componentInstance[name]
	if !ok {
		c, err := h.createComponent(ctx, name)
		if err != nil {
			fmt.Printf("create component %s failed\n", name)
			return nil, false
		}
		h.componentInstance[name] = c
	}
	return c, true
}

func (h *HostImpl) createComponent(ctx context.Context, name string) (comstate.Store, error) {
	spec, ok := h.componentSpec[name]
	if !ok {
		return nil, errors.New("component not found")
	}
	s, err := registery.DefaultRegistry.Create(name, "v1", "state."+name)
	if err != nil {
		return nil, err
	}
	baseMetadata := toBaseMetadata(spec)
	err = s.Init(ctx, comstate.Metadata{Base: baseMetadata})
	if err != nil {
		return nil, err
	}
	return s, nil
}

func convertMetadataItemsToProperties(items []MetadataItem) map[string]string {
	properties := map[string]string{}
	for _, c := range items {
		val := c.Value
		properties[c.Name] = val
	}
	return properties
}

func toBaseMetadata(c Component) contribMetadata.Base {
	return contribMetadata.Base{
		Properties: convertMetadataItemsToProperties(c.Spec.Metadata),
		Name:       c.Name,
	}
}

// todo check why set signature is different with get signature

func (h *HostImpl) Get(ctx context.Context, m api.Module, namePtr, nameLen, keyPtr, kenLen, metadataOption, metadataPtr, metadataLen, consistency, resultPtr uint32) {
	// read name
	name, ok := m.Memory().Read(namePtr, nameLen)
	if !ok {
		handleGetError(ctx, m, resultPtr, errors.New("read name failed"))
		return
	}

	// read key
	key, ok := m.Memory().Read(keyPtr, kenLen)
	if !ok {
		handleGetError(ctx, m, resultPtr, errors.New("read key failed"))
		return
	}

	// decode metadata
	var metadata map[string]string
	if metadataOption == 1 {
		data, ok := m.Memory().Read(metadataPtr, metadataLen*16)
		if !ok {
			handleGetError(ctx, m, resultPtr, errors.New("read metadata memory failed"))
			return
		}
		var err error
		metadata, err = decodeMetadata(ctx, m, data, metadataLen)
		if err != nil {
			handleGetError(ctx, m, resultPtr, err)
			return
		}
	}
	c, ok := h.getComponent(ctx, string(name))
	if !ok {
		handleGetError(ctx, m, resultPtr, errors.New("get component failed"))
		return
	}
	resp, err := c.Get(ctx, &comstate.GetRequest{
		Key:      string(key),
		Metadata: metadata,
		Options:  comstate.GetStateOption{Consistency: stateConsistencyToString(consistency)},
	})
	if err != nil {
		handleGetError(ctx, m, resultPtr, err)
		return
	}

}

// todo free memory
func marshalGetResponse(ctx context.Context, m api.Module, resp *comstate.GetResponse) []byte {
	resbytes := make([]byte, GetLen)
	binary.LittleEndian.PutUint32(resbytes[0:4], 0) // success
	if resp.ETag != nil {
		binary.LittleEndian.PutUint32(resbytes[12:16], 1) // has etag
		mp, err := common.Malloc(ctx, m, uint32(len(*resp.ETag)))
		if err != nil {
			fmt.Println("malloc failed")
			return nil
		}
		binary.LittleEndian.PutUint32(resbytes[16:20], mp)
		binary.LittleEndian.PutUint32(resbytes[20:24], uint32(len(*resp.ETag)))
	} else {
		binary.LittleEndian.PutUint32(resbytes[12:16], 0) // no etag
	}
	// handle metadata
	if len(resp.Metadata) > 0 {
		binary.LittleEndian.PutUint32(resbytes[24:28], 1) // has metadata
		mp, err := common.Malloc(ctx, m, uint32(len(resp.Metadata)*16))
		if err != nil {
			fmt.Println("malloc failed")
			return nil
		}
		binary.LittleEndian.PutUint32(resbytes[28:32], mp)
		binary.LittleEndian.PutUint32(resbytes[32:36], uint32(len(resp.Metadata)))
		mt, err := encodeMetadata(ctx, m, resp.Metadata)
		if err != nil {
			fmt.Println("encode metadata failed")
			return nil
		}
		m.Memory().Write(mp, mt)
	} else {
		binary.LittleEndian.PutUint32(resbytes[24:28], 0) // no metadata
	}

	// handle content type
	if resp.ContentType != nil {
		binary.LittleEndian.PutUint32(resbytes[36:40], 1) // has content type
		mp, err := common.Malloc(ctx, m, uint32(len(*resp.ContentType)))
		if err != nil {
			fmt.Println("malloc failed")
			return nil
		}
		binary.LittleEndian.PutUint32(resbytes[40:44], mp)
		binary.LittleEndian.PutUint32(resbytes[44:48], uint32(len(*resp.ContentType)))
	} else {
		binary.LittleEndian.PutUint32(resbytes[36:40], 0) // no content type
	}
	return resbytes
}

func handleGetError(ctx context.Context, m api.Module, ptr uint32, err error) {
	resbytes := make([]byte, GetLen)
	binary.LittleEndian.PutUint32(resbytes[0:4], 1) // err
	msg := err.Error()
	mp, err := common.Malloc(ctx, m, uint32(len(msg)))
	if err != nil {
		fmt.Println("malloc failed")
		return
	}
	m.Memory().Write(mp, []byte(msg))
	binary.LittleEndian.PutUint32(resbytes[4:8], mp)
	binary.LittleEndian.PutUint32(resbytes[8:12], uint32(len(msg)))
	m.Memory().Write(ptr, resbytes)
}

func decodeMetadata(ctx context.Context, mod api.Module, data []byte, itemCount uint32) (map[string]string, error) {
	fields := make(map[string]string)
	for i := uint32(0); i < itemCount; i++ {
		key_ptr := binary.LittleEndian.Uint32(data[i*16 : i*16+4])
		key_len := binary.LittleEndian.Uint32(data[i*16+4 : i*16+8])
		key, ok := common.ReadString(mod, key_ptr, key_len)
		if !ok {
			fmt.Println("Error reading key")
			return nil, errors.New("error reading key")
		}
		val_ptr := binary.LittleEndian.Uint32(data[i*16+8 : i*16+12])
		val_len := binary.LittleEndian.Uint32(data[i*16+12 : i*16+16])
		val, ok := common.ReadString(mod, val_ptr, val_len)
		if !ok {
			fmt.Println("Error reading value")
			return nil, errors.New("error reading value")
		}
		//if _, found := fields[key]; !found {
		//	fields[key] = []string{}
		//}
		fields[key] = val
	}
	return fields, nil
}

func encodeMetadata(ctx context.Context, mod api.Module, metadata map[string]string) ([]byte, error) {
	var metadataBytes []byte = make([]byte, 16*len(metadata))
	var i int
	for k, v := range metadata {
		keyPtr, err := common.Malloc(ctx, mod, uint32(len(k)))
		if err != nil {
			return nil, err
		}
		valPtr, err := common.Malloc(ctx, mod, uint32(len(v)))
		if err != nil {
			return nil, err
		}
		binary.LittleEndian.PutUint32(metadataBytes[i*16+0:i*16+4], keyPtr)
		binary.LittleEndian.PutUint32(metadataBytes[i*16+4:i*16+8], uint32(len(k)))
		binary.LittleEndian.PutUint32(metadataBytes[i*16+8:i*16+12], valPtr)
		binary.LittleEndian.PutUint32(metadataBytes[i*16+12:i*16+16], uint32(len(v)))
		i++
	}
	return metadataBytes, nil
}

func (h *HostImpl) Set(ctx context.Context, m api.Module, inPtr, outPtr uint32) {

}

func (h *HostImpl) Delete(name string, req state.DaprStateStateTypesDeleteRequest) state.Result[uint32, string] {
	//TODO implement me
	panic("implement me")
}
