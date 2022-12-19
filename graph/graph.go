// Developed by Kaiser925 on 2022/12/19.
// Lasted modified 2022/12/19.
// Copyright (c) 2022.  All rights reserved
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graph

// ValidPath returns whether there is path between source and destination.
func ValidPath(n int, edges [][]int, source int, destination int) bool {
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if p[x] != x {
			p[x] = find(p[x])
		}
		return p[x]
	}
	for _, e := range edges {
		p[find(e[0])] = find(e[1])
	}
	return find(source) == find(destination)
}
