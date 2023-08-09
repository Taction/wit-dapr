


.PHONY: wit
wit:
	@echo "Building wit..."
	wit-bindgen tiny-go --out-dir dapr/state wit/state
	wit-bindgen tiny-go --out-dir dapr/http wit/http
	#wit-bindgen tiny-go --out-dir dapr/keyvalue wit/keyvalue