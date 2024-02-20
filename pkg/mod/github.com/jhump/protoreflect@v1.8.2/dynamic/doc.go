// Package dynamic provides an implementation for a dynamic protobuf message.
//
// The dynamic message is essentially a message descriptor along with a map of
// tag numbers to values. It has a broad API for interacting with the message,
// including inspection and modification. Generally, most operations have two
// forms: a regular method that panics on bad input or error and a "Try" form
// of the method that will instead return an error.
//
// A dynamic message can optionally be constructed with a MessageFactory. The
// MessageFactory has various registries that may be used by the dynamic message,
// such as during de-serialization. The message factory is "inherited" by any
// other dynamic messages created, such as nested messages that are created
// during de-serialization. Similarly, any dynamic message created using
// MessageFactory.NewMessage will be associated with that factory, which in turn
// will be used to create other messages or parse extension fields during
// de-serialization.
//
//
// Field Types
//
// The types of values expected by setters and returned by getters are the
// same as protoc generates for scalar fields. For repeated fields, there are
// methods for getting and setting values at a particular index or for adding
// an element. Similarly, for map fields, there are methods for getting and
// setting values for a particular key.
//
// If you use GetField for a repeated field, it will return a copy of all
// elements as a slice []interface{}. Similarly, using GetField for a map field
// will return a copy of all mappings as a map[interface{}]interface{}. You can
// also use SetField to supply an entire slice or map for repeated or map fields.
// The slice need not be []interface{} but can actually be typed according to
// the field's expected type. For example, a repeated uint64 field can be set
// using a slice of type []uint64.
//
// Descriptors for map fields describe them as repeated fields with a nested
// message type. The nested message type is a special generated type that
// represents a single mapping: key and value pair. The dynamic message has some
// special affordances for this representation. For example, you can use
// SetField to set a map field using a slice of these entry messages. Internally,
// the slice of entries will be converted to an actual map. Similarly, you can
// use AddRepeatedField with an entry message to add (or overwrite) a mapping.
// However, you cannot use GetRepeatedField or SetRepeatedField to modify maps,
// since those take numeric index arguments which are not relevant to maps
// (since maps in Go have no defined ordering).
//
// When setting field values in dynamic messages, the type-checking is lenient
// in that it accepts any named type with the right kind. So a string field can
// be assigned to any type that is defined as a string. Enum fields require
// int32 values (or any type that is defined as an int32).
//
// Unlike normal use of numeric values in Go, values will be automatically
// widened when assigned. So, for example, an int64 field can be set using an
// int32 value since it can be safely widened without truncation or loss of
// precision. Similar goes for uint32 values being converted to uint64 and
// float32 being converted to float64. Narrowing conversions are not done,
// however. Also, unsigned values will never be automatically converted to
// signed (and vice versa), and floating point values will never be
// automatically converted to integral values (and vice versa). Since the bit
// width of int and uint fields is allowed to be platform dependent, but will
// always be less than or equal to 64, they can only be used as values for
// int64 and uint64 fields, respectively. They cannot be used to set int32 or
// uint32 fields, which includes enums fields.
//
// Fields whose type is a nested message can have values set to either other
// dynamic messages or generated messages (e.g. pointers to structs generated by
// protoc). Getting a value for such a field will return the actual type it is
// set to (e.g. either a dynamic message or a generated message). If the value
// is not set and the message uses proto2 syntax, the default message returned
// will be whatever is returned by the dynamic message's MessageFactory (if the
// dynamic message was not created with a factory, it will use the logic of the
// zero value factory). In most typical cases, it will return a dynamic message,
// but if the factory is configured with a KnownTypeRegistry, or if the field's
// type is a well-known type, it will return a zero value generated message.
//
//
// Unrecognized Fields
//
// Unrecognized fields are preserved by the dynamic message when unmarshaling
// from the standard binary format. If the message's MessageFactory was
// configured with an ExtensionRegistry, it will be used to identify and parse
// extension fields for the message.
//
// Unrecognized fields can dynamically become recognized fields if the
// application attempts to retrieve an unrecognized field's value using a
// FieldDescriptor. In this case, the given FieldDescriptor is used to parse the
// unknown field and move the parsed value into the message's set of known
// fields. This behavior is most suited to the use of extensions, where an
// ExtensionRegistry is not setup with all known extensions ahead of time. But
// it can even happen for non-extension fields! Here's an example scenario where
// a non-extension field can initially be unknown and become known:
//
//   1. A dynamic message is created with a descriptor, A, and then
//      de-serialized from a stream of bytes. The stream includes an
//      unrecognized tag T. The message will include tag T in its unrecognized
//      field set.
//   2. Another call site retrieves a newer descriptor, A', which includes a
//      newly added field with tag T.
//   3. That other call site then uses a FieldDescriptor to access the value of
//      the new field. This will cause the dynamic message to parse the bytes
//      for the unknown tag T and store them as a known field.
//   4. Subsequent operations for tag T, including setting the field using only
//      tag number or de-serializing a stream that includes tag T, will operate
//      as if that tag were part of the original descriptor, A.
//
//
// Compatibility
//
// In addition to implementing the proto.Message interface, the included
// Message type also provides an XXX_MessageName() method, so it can work with
// proto.MessageName. And it provides a Descriptor() method that behaves just
// like the method of the same signature in messages generated by protoc.
// Because of this, it is actually compatible with proto.Message in many (though
// not all) contexts. In particular, it is compatible with proto.Marshal and
// proto.Unmarshal for serializing and de-serializing messages.
//
// The dynamic message supports binary and text marshaling, using protobuf's
// well-defined binary format and the same text format that protoc-generated
// types use. It also supports JSON serialization/de-serialization by
// implementing the json.Marshaler and json.Unmarshaler interfaces. And dynamic
// messages can safely be used with the jsonpb package for JSON serialization
// and de-serialization.
//
// In addition to implementing the proto.Message interface and numerous related
// methods, it also provides inter-op with generated messages via conversion.
// The ConvertTo, ConvertFrom, MergeInto, and MergeFrom methods copy message
// contents from a dynamic message to a generated message and vice versa.
//
// When copying from a generated message into a dynamic message, if the
// generated message contains fields unknown to the dynamic message (e.g. not
// present in the descriptor used to create the dynamic message), these fields
// become known to the dynamic message (as per behavior described above in
// "Unrecognized Fields"). If the generated message has unrecognized fields of
// its own, including unrecognized extensions, they are preserved in the dynamic
// message. It is possible that the dynamic message knows about fields that the
// generated message did not, like if it has a different version of the
// descriptor or its MessageFactory has an ExtensionRegistry that knows about
// different extensions than were linked into the program. In this case, these
// unrecognized fields in the generated message will be known fields in the
// dynamic message.
//
// Similarly, when copying from a dynamic message into a generated message, if
// the dynamic message has unrecognized fields they can be preserved in the
// generated message (currently only for syntax proto2 since proto3 generated
// messages do not preserve unrecognized fields). If the generated message knows
// about fields that the dynamic message does not, these unrecognized fields may
// become known fields in the generated message.
//
//
// Registries
//
// This package also contains a couple of registries, for managing known types
// and descriptors.
//
// The KnownTypeRegistry allows de-serialization of a dynamic message to use
// generated message types, instead of dynamic messages, for some kinds of
// nested message fields. This is particularly useful for working with proto
// messages that have special encodings as JSON (e.g. the well-known types),
// since the dynamic message does not try to handle these special cases in its
// JSON marshaling facilities.
//
// The ExtensionRegistry allows for recognizing and parsing extensions fields
// (for proto2 messages).
package dynamic
