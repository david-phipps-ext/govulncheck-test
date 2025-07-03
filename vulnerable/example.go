package vulnerable

import (
    "gopkg.in/yaml.v2"
)

// Example function that uses a potentially vulnerable dependency
func ParseYAML(data []byte) (map[string]interface{}, error) {
    var result map[string]interface{}
    err := yaml.Unmarshal(data, &result)
    return result, err
}
