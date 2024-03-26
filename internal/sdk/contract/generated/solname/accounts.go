// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package solname

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type AdminConfig struct {
	PriceFeed        ag_solanago.PublicKey
	ChainlinkProgram ag_solanago.PublicKey
}

var AdminConfigDiscriminator = [8]byte{156, 10, 79, 161, 71, 9, 62, 77}

func (obj AdminConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(AdminConfigDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `PriceFeed` param:
	err = encoder.Encode(obj.PriceFeed)
	if err != nil {
		return err
	}
	// Serialize `ChainlinkProgram` param:
	err = encoder.Encode(obj.ChainlinkProgram)
	if err != nil {
		return err
	}
	return nil
}

func (obj *AdminConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(AdminConfigDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[156 10 79 161 71 9 62 77]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `PriceFeed`:
	err = decoder.Decode(&obj.PriceFeed)
	if err != nil {
		return err
	}
	// Deserialize `ChainlinkProgram`:
	err = decoder.Decode(&obj.ChainlinkProgram)
	if err != nil {
		return err
	}
	return nil
}

type Domain struct {
	Authority   ag_solanago.PublicKey
	Name        string
	ExpiresAt   int64
	Profile     Profile
	Socials     SocialMedia
	Wallets     Wallets
	TextRecords []TextRecord
	Subdomains  []SubDomain
	Mint        ag_solanago.PublicKey
	NftToken    ag_solanago.PublicKey
	Metadata    ag_solanago.PublicKey
}

var DomainDiscriminator = [8]byte{167, 191, 231, 63, 146, 41, 115, 27}

func (obj Domain) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(DomainDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Authority` param:
	err = encoder.Encode(obj.Authority)
	if err != nil {
		return err
	}
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	// Serialize `ExpiresAt` param:
	err = encoder.Encode(obj.ExpiresAt)
	if err != nil {
		return err
	}
	// Serialize `Profile` param:
	err = encoder.Encode(obj.Profile)
	if err != nil {
		return err
	}
	// Serialize `Socials` param:
	err = encoder.Encode(obj.Socials)
	if err != nil {
		return err
	}
	// Serialize `Wallets` param:
	err = encoder.Encode(obj.Wallets)
	if err != nil {
		return err
	}
	// Serialize `TextRecords` param:
	err = encoder.Encode(obj.TextRecords)
	if err != nil {
		return err
	}
	// Serialize `Subdomains` param:
	err = encoder.Encode(obj.Subdomains)
	if err != nil {
		return err
	}
	// Serialize `Mint` param:
	err = encoder.Encode(obj.Mint)
	if err != nil {
		return err
	}
	// Serialize `NftToken` param:
	err = encoder.Encode(obj.NftToken)
	if err != nil {
		return err
	}
	// Serialize `Metadata` param:
	err = encoder.Encode(obj.Metadata)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Domain) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(DomainDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[167 191 231 63 146 41 115 27]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Authority`:
	err = decoder.Decode(&obj.Authority)
	if err != nil {
		return err
	}
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	// Deserialize `ExpiresAt`:
	err = decoder.Decode(&obj.ExpiresAt)
	if err != nil {
		return err
	}
	// Deserialize `Profile`:
	err = decoder.Decode(&obj.Profile)
	if err != nil {
		return err
	}
	// Deserialize `Socials`:
	err = decoder.Decode(&obj.Socials)
	if err != nil {
		return err
	}
	// Deserialize `Wallets`:
	err = decoder.Decode(&obj.Wallets)
	if err != nil {
		return err
	}
	// Deserialize `TextRecords`:
	err = decoder.Decode(&obj.TextRecords)
	if err != nil {
		return err
	}
	// Deserialize `Subdomains`:
	err = decoder.Decode(&obj.Subdomains)
	if err != nil {
		return err
	}
	// Deserialize `Mint`:
	err = decoder.Decode(&obj.Mint)
	if err != nil {
		return err
	}
	// Deserialize `NftToken`:
	err = decoder.Decode(&obj.NftToken)
	if err != nil {
		return err
	}
	// Deserialize `Metadata`:
	err = decoder.Decode(&obj.Metadata)
	if err != nil {
		return err
	}
	return nil
}

type DnsState struct {
	Admin             ag_solanago.PublicKey
	PrimaryDomains    []ag_solanago.PublicKey
	AllowedTopDomains []string
}

var DnsStateDiscriminator = [8]byte{212, 184, 33, 18, 41, 35, 205, 88}

func (obj DnsState) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(DnsStateDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Admin` param:
	err = encoder.Encode(obj.Admin)
	if err != nil {
		return err
	}
	// Serialize `PrimaryDomains` param:
	err = encoder.Encode(obj.PrimaryDomains)
	if err != nil {
		return err
	}
	// Serialize `AllowedTopDomains` param:
	err = encoder.Encode(obj.AllowedTopDomains)
	if err != nil {
		return err
	}
	return nil
}

func (obj *DnsState) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(DnsStateDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[212 184 33 18 41 35 205 88]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Admin`:
	err = decoder.Decode(&obj.Admin)
	if err != nil {
		return err
	}
	// Deserialize `PrimaryDomains`:
	err = decoder.Decode(&obj.PrimaryDomains)
	if err != nil {
		return err
	}
	// Deserialize `AllowedTopDomains`:
	err = decoder.Decode(&obj.AllowedTopDomains)
	if err != nil {
		return err
	}
	return nil
}

type PrimaryDomain struct {
	Name string
}

var PrimaryDomainDiscriminator = [8]byte{231, 255, 61, 63, 142, 184, 254, 42}

func (obj PrimaryDomain) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(PrimaryDomainDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	return nil
}

func (obj *PrimaryDomain) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(PrimaryDomainDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[231 255 61 63 142 184 254 42]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	return nil
}

type Decimal struct {
	Value    ag_binary.Int128
	Decimals uint32
}

var DecimalDiscriminator = [8]byte{138, 41, 153, 49, 98, 29, 40, 56}

func (obj Decimal) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(DecimalDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Value` param:
	err = encoder.Encode(obj.Value)
	if err != nil {
		return err
	}
	// Serialize `Decimals` param:
	err = encoder.Encode(obj.Decimals)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Decimal) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(DecimalDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[138 41 153 49 98 29 40 56]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Value`:
	err = decoder.Decode(&obj.Value)
	if err != nil {
		return err
	}
	// Deserialize `Decimals`:
	err = decoder.Decode(&obj.Decimals)
	if err != nil {
		return err
	}
	return nil
}