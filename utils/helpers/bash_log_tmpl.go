package helpers

const cdxBashLogTemplate = `#!/usr/bin/env bash

## <editor-fold defaultstate="collapsed" desc="Convert case">
function mri_camel_to_snake() {
  printf '%s' "${1:-}" | sed -E 's/([A-Z])/_\L\1/g' | sed 's/^_//' || return 1
} 
function mri_camel_to_kebab() {
  printf '%s' "${1:-}" | sed -E 's/([A-Z])/-\L\1/g' | sed 's/^-//' || return 1
}
function mri_convert_case() { 
  local text="$1"
  local target_case="$2"

  case "$target_case" in
    snake) printf '%s' "$text" | sed -E 's/([A-Z])/_\L\1/g' | sed 's/^_//' ;;
    kebab) printf '%s' "$text" | sed -E 's/([A-Z])/-\L\1/g' | sed 's/^-//' ;;
    camel) printf '%s' "$text" | awk -F'[_-]' '{for (i=1; i<=NF; i++) $i=toupper(substr($i,1,1)) substr($i,2)}1' OFS='' ;;
    *) printf '%s\n' "Invalid case. Use 'snake', 'kebab', or 'camel'." && return 1 ;;
  esac
}
## </editor-fold>

## <editor-fold defaultstate="collapsed" desc="Colors">
function mri_get_colors(){
  if [ -n "${NO_COLOR:-}" ] || [ -n "${ANSI_COLORS_DISABLED:-}" ]; then
    export nocolor="" bold="" nobold="" underline="" nounderline="" red="" green="" yellow="" blue="" magenta="" cyan=""
    return 0
  fi

  if ! tput colors &>/dev/null || [ "$(tput colors)" -lt 8 ]; then
    export nocolor="" bold="" nobold="" underline="" nounderline="" red="" green="" yellow="" blue="" magenta="" cyan=""
    return 0
  fi

  export nocolor="\033[0m" bold="\033[1m" nobold="\033[22m"
  export underline="\033[4m" nounderline="\033[24m"
  export red="\033[31m" green="\033[32m" yellow="\033[33m" blue="\033[34m" magenta="\033[35m" cyan="\033[36m"
}
## </editor-fold>

## <editor-fold defaultstate="collapsed" desc="Log">
function mri_log() {
  local log_type="${1:-info}"
  shift
  local log_msg="${*:-}"
  local color=""

  case "$log_type" in
    error|fatal) color="${bold}${red}" ;;
    warn|alert) color="${bold}${yellow}" ;;
    info|debug) color="${bold}${blue}" ;;
    success) color="${bold}${green}" ;;
    *) color="${nocolor}" ;;
  esac

  printf '%b[%s]%b %s\n' "$color" "$log_type" "$nocolor" "$log_msg"
}
## </editor-fold>

## <editor-fold defaultstate="collapsed" desc="Traps">
function mri_set_trap(){
  trap 'mri_handle_exit $? ${LINENO:-}' ERR
  trap 'mri_handle_exit $? ${LINENO:-}' EXIT HUP INT QUIT ABRT ALRM TERM
}

function mri_handle_exit() {
  local exit_code="${1}"
  local line_number="${2:-0}"

  if [ "$exit_code" -eq 0 ]; then
    exit 0
  fi

  mri_log error "Erro ocorrido na linha: ${line_number}"
  exit "$exit_code"
}
## </editor-fold>

mri_get_colors

## Only shows the message if the script is executed directly
if [ "${BASH_SOURCE[0]}" = "${0}" ]; then
  if [[ ! -f $HOME/.cache/.mri_install_log ]]; then
	if [[ ! -f $HOME/.cache/.mri_usage_warning ]]; then
		mri_log warn "💡 (This script should not be executed directly)"
		mri_log warn "💡 (It is intended to be sourced by another script)"
		mri_log warn "💡 (Example: 'source kubex_helpers.sh')"
		touch $HOME/.cache/.mri_usage_warning
	fi
	mri_log info "✅ Kubex Helpers loaded successfully."
  fi
fi

`
