package helpers

const cdxYesNoQyestionTemplate = `#!/bin/bash

function _mri_yes_no_question() {
  source $HOME/.mori_logging_env

  local _question="${1:-$2}"
  local _default_answer="${2:-}"
  local _timeout="${3:-5}"
  local _answer=""
  local _counter=0
  
  while [[ ! "$_answer" =~ ^[yYnN]$ ]]; do
    if [ "$_counter" -gt 3 ]; then
      mri_log error "Maximum number of attempts reached."
      printf '%s\n' "Please try again later."
      return 1 
    else
      _counter=$((_counter + 1))
    fi
    read -rt "${_timeout}" -rp "${_question} " -n 1 _answer || _answer="${_default_answer:-}"
    if [ -z "${_answer}" ]; then
      if [ -n "${_default_answer}" ]; then
        _answer="${_default_answer}"
      fi
    fi
  done
  if [[ "${_answer}" =~ ^[yY]$ ]]; then
    # printf '%s\n' "y"
    return 0
  else
    # printf '%s\n' "n"
    return 1
  fi
}

_mri_yes_no_question "$@"

exit $?

`
