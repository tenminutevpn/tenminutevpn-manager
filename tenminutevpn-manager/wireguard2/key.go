package wireguard2

import "golang.zx2c4.com/wireguard/wgctrl/wgtypes"

type Key wgtypes.Key

func (k Key) PublicKey() Key {
	return Key(wgtypes.Key(k).PublicKey())
}

func (k Key) String() string {
	return wgtypes.Key(k).String()
}

func (k Key) MarshalYAML() (interface{}, error) {
	return k.String(), nil
}

func (k *Key) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var key string
	if err := unmarshal(&key); err != nil {
		return err
	}

	if key == "" {
		return nil
	}

	wireguardKey, err := wgtypes.ParseKey(key)
	if err != nil {
		return err
	}

	*k = Key(wireguardKey)
	return nil
}

func GenerateKey() (Key, error) {
	key, err := wgtypes.GenerateKey()
	if err != nil {
		return Key{}, err
	}
	return Key(key), nil

}
