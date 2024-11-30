echo "(Loading) Custom Dockerfile shell profile settings" >&2

export PATH="$HOME/go/bin:$PATH"
export TENV_AUTO_INSTALL=true

alias typos="$HOME/.cache/pre-commit/repoyq9jys2r/py_env-python3.10/bin/typos"

# CI Workflow specific variables
if [ -n "${GITHUB_WORKFLOW}" ]; then
	echo "CI workflow detected."

	# Disable pre-commit hook that blocks files with autogen in their path
	SKIP="no-autogen-on-dev-machines"
fi
