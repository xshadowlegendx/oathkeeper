#!/usr/bin/env bash
# Copyright 2021 Ory GmbH
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


set -euo pipefail

cd "$( dirname "${BASH_SOURCE[0]}" )/.."

toformat=$(goimports -l $(go list -f {{.Dir}} ./... | grep -v vendor | grep -v 'oathkeeper$'))
[ -z "$toformat" ] && echo "All files are formatted correctly"
[ -n "$toformat" ] && echo "Please use \`goimports\` to format the following files:" && echo $toformat && exit 1

exit 0
