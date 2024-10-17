#!/usr/bin/env bash

HOST="vyos.local"

SCRIPT_DIR=$(dirname -- "$(readlink -f -- "$BASH_SOURCE")")
KEY="MySuperSecretKey"

function main() {
	local cmd="$1"
	shift

	case "$cmd" in
	get)
		api_get "$@"
		;;
	set)
		api_set "$@"
		;;
	has)
		api_has "$@"
		;;
	*)
		echo "err"
		exit 1
		;;
	esac
}

function api_get() {
	curl -v -k --location --request POST "https://$HOST/retrieve" \
		--form key="$KEY" \
		--form data='{"op": "showConfig", "path": []}'
}

function api_set() {
	curl -k --location --request POST "https://$HOST/configure" \
		--form key="$KEY" \
		--form data=$(
			cat <<-EOL
				[{"op":"set","path":["service","ntp","allow-client","address","fc00::/7"]},{"op":"set","path":["service","ntp","allow-client","address","10.0.0.0/8"]},{"op":"set","path":["service","ntp","allow-client","address","fc00::/7"]},{"op":"set","path":["service","ntp","allow-client","address","fc00::/7"]},{"op":"set","path":["service","ntp","allow-client","address","fc00::/7"]},{"op":"set","path":["service","ntp","allow-client","address","fc00::/7"]},{"op":"set","path":["service","ntp","allow-client","address","fc00::/7"]},{"op":"set","path":["service","ntp","allow-client","address","fc00::/7"]}]
			EOL
		) | jq
}

function api_has() {
	curl -v -k --location --request POST "https://$HOST/retrieve" \
		--form key="$KEY" \
		--form data=$(
			cat <<-EOL
				[{"op":"exists","path":["system","conntrack","tcp"]}]
			EOL
		) | jq
}

main "$@"
