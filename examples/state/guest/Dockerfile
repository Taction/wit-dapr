FROM scratch
COPY ./main.wasm /main.wasm
COPY ./config.yaml /config.yaml
COPY ./redis-state.yaml /resource-path/redis-state.yaml
ENTRYPOINT [ "main.wasm" ]