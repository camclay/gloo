// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/failover.proto

package v1

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"
	"strconv"

	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
	"github.com/solo-io/protoc-gen-ext/pkg/hasher/hashstructure"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = new(hash.Hash64)
	_ = fnv.New64
	_ = strconv.Itoa
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Failover) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1.Failover")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("PrioritizedLocalities")); err != nil {
		return 0, err
	}
	for i, v := range m.GetPrioritizedLocalities() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("v")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	if h, ok := interface{}(m.GetPolicy()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Policy")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetPolicy(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Policy")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *LocalityLbEndpoints) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1.LocalityLbEndpoints")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetLocality()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Locality")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetLocality(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Locality")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("LbEndpoints")); err != nil {
		return 0, err
	}
	for i, v := range m.GetLbEndpoints() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("v")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	if h, ok := interface{}(m.GetLoadBalancingWeight()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("LoadBalancingWeight")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetLoadBalancingWeight(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("LoadBalancingWeight")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *LbEndpoint) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1.LbEndpoint")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Address")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetAddress())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Port")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetPort())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetHealthCheckConfig()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("HealthCheckConfig")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetHealthCheckConfig(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("HealthCheckConfig")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetUpstreamSslConfig()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("UpstreamSslConfig")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetUpstreamSslConfig(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("UpstreamSslConfig")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetLoadBalancingWeight()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("LoadBalancingWeight")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetLoadBalancingWeight(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("LoadBalancingWeight")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetMetadata() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("v")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("v")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
			}

			if _, err = innerHash.Write([]byte("k")); err != nil {
				return 0, err
			}
			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Locality) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1.Locality")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Region")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetRegion())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Zone")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetZone())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("SubZone")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetSubZone())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Failover_PrioritizedLocality) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1.Failover_PrioritizedLocality")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("LocalityEndpoints")); err != nil {
		return 0, err
	}
	for i, v := range m.GetLocalityEndpoints() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("v")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Failover_Policy) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1.Failover_Policy")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetOverprovisioningFactor()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("OverprovisioningFactor")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetOverprovisioningFactor(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("OverprovisioningFactor")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *LbEndpoint_HealthCheckConfig) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1.LbEndpoint_HealthCheckConfig")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("PortValue")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetPortValue())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Hostname")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetHostname())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Path")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetPath())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Method")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetMethod())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
