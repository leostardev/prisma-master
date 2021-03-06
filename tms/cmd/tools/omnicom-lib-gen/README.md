# Introduction 

Libomnicom is a library written in golang to decode omnicom solar data messages in byte arrays into golang structures and decode them from golang structures in turn.

# Parsing

Libomnicom uses a method Parse in the omnicom package to decode omnicom solar messages

func Parse(byteList []byte) (Omnicom, error)

### Arguments:
- byteList: the array of bytes of omnicom solar messages

### Return Values

- Omnicom: Message interface containing all omnicom messages types defined in golang structures
- Error: return an error if the input bytes have no valid corresponding messages.

# Encoding

Libomnicom has a method Encode for each message type to encode it to messages in byte arrays

func (msg *<Msg sturcture>) Encode() ([]byte, error)

### Return Values:
- []byte: byte arrays of omnicom messages
- Error: return an error if the message to be encoded is not initialized correctly

### Message structure list

- HPR: History position report (0x01), Iridium/3G
- SPR: Single position report (0x06), Iridium/3G
- AR: Alert Report (0x02), Iridium/3G
- GP: Global parameters(0x03), Iridium/3G
- AUP: API url parameters (0x08), Iridium/3G
- GA: Geofencing ack(0x04), Iridium/3G
- BM: Binary Message (0x07), Iridium/3G
- ABM: Ack binary message (0x09), Iridium/3G
- TMA: Test Mode Ack (0x30), Iridium/3G
- AA: Ack assistance (0x45), Iridium/3G
- RMH: Request Message History (0x31), Iridium/3G
- UIC: Unit interval change (0x32), Iridium/3G
- RSM: Request a specific message (0x33), Iridium/3G
- UGP: Update global parameters (0x34), Iridium/3G
- UAUP: Update API url parameters (0x3A), Iridium/3G
- UG_Polygon: Upload Geofence (0x35) for polygon shape, Iridium/3G
- UG_Circle: Upload Geofence (0x35) for circle shape, Iridium/3G
- GBMN: 3G binary message notification (0x36), Iridium
- DGF: Delete Geofence (0x37), Iridium/3G
- BM_StoV: Binary message (0x38) from server to vessels, Iridium/3G
- UF: Update Firmware (0x39), 3G
- SPM: Split Packet message (0x40), Iridium
- SDR: Split diagnostic request (0x41), Iridium
- SMDR: Split message diagnostic response (0x05), Iridium

# Extensibility

Libomnicom code is generated by gen.go according to the Messages.json. More message types can be added in the Messages.json file. Then run ./gen with the json file in the same directory to generate lib code in ./omnicom.

## Structure of Messages.json

Each message is defined as a Message struct in the Messages.json. Each Message struct may have mutiple Field structs. Each Field which is not in type byte, uint32, uint64, float32, []rune, []byte or string may have mutiple sub Fields.

Fields that have "Extention_Bit" as a prefix are indicating that the following field in the message is optional. The generator uses the existance of extention bit fields to determine the the proper way to generate encoders/decoders for optional fields. 

### Message struct

```go
type Message struct {
	Name      string
	Type      string
	Explain   string
	HeaderStr string
	Ident     string
	Index     int
	Fields    []Field
	Size      int
}
```

- Name: The name of this Message. It is not necessary to be indicated in json file. It can be determined automatically according to it type by gen.go.
- Type: The type of this Message. Usually set as the acronyms of messages' introduction as shown in the Message structure list.
- Explain: The introduction of this Message. It is printed as a comment at the head of the lib code of this Message.
- HeaderStr: Binary string of the header byte of this Message. For example,  header 0x06 has a HeaderStr "00000110". It is used to find a corresponding parser when parsing a input byte array.
- Ident: Used as a identity to differentiate Messages which have a same header. For example, both UG_Polygon and UG_Circle has a same header 0x35. UG_Polygon's Shape is 0 and UG_Circle's Shape is 1. So UG_Polygon's Ident is "0" and UG_Circle's Ident is "1". In most Messages, it need not to be set.
- Index: Used as the index of the identity in the binary strings. UG_Polygon and UG_Circle's Index is 275. Because the 276th digit of them stands for the field Shape.
- Fields: All Fields of this Message.
- Size: The length of the binary strings of fixed length Messages. For the flexible length Messags, this is the maximum of their binary strings' length.

### Field struct 

```go
type Field struct {
	Name      string
	Repeated  int
	Type      string
	Fields    []Field
	Size      int
	Reference string
	Offset    string
	Scale     string
}
```

- Name: The name of this Field. Make sure its first letter is in uppercase.
- Repeated: If its Message has multiple instances of this Field. Its Repeated is defined as the maximum of the repeated times. For example the Field Data_Report's Repeated is 13 in HPR. Its default value is 1.
- Type: The type of this Field. It can be byte, uitn32, uint64, float32, []rune, []byte, string or a self-defined type. Only Field Header of each Message can be in type byte. If it is a self-defined one, this Field has sub-fields which are defined in Fields. Otherwise the Fields array is empty. If a self-defined Field type is defined in this json file already, its Fields need not to be declared again if it occurs in another Message or Field because gen.go has known it.
- Fields: The collection of its sub-fields if it is a self-defined one.
- Size: The length of its binary string.
- Reference: Two kinds of Field may have References. One is Field in type string. Its reference is another Field in its Message which determines its length in byte. For instance, the Msg_Data_Part in SPM has its Reference set as "Length_Msg_Data_Part_in_Byte". The other is self-defined Field. Its Reference is another Field in its Message which determines its number of times to repeat. For instance, the Data_Report in HPR has is Reference set as "Count_Data_Reports_in_this_Msg".
- Offset: The offset to get its actual value from its coded one. Its default is "0".
- Scale: The scale to get its actual value from its coded one. Its default is "1".

## Generate lib code

Make sure the file Messages.json, gen.go and a folder ./omnicom are in the same directory. Use this command to generate the lib code in ./Omnicom.

go build -o gen gen.go