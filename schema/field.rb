require 'string_inflection'
using StringInflection

class Field
  attr_reader \
    :datetime_layout,
    :deserializer,
    :identifier,
    :go_type,
    :go_field_name,
    :go_field_tag,
    :description,
    :hash

  GO_TYPES = %w(
    string
    bool
    time.Time
    int
    []string
    float64
  ).freeze

  def initialize(hash)
    return if hash.nil?

    @hash = hash
    @datetime_layout = hash[:datetimeLayout] || 'time.RFC3339Nano'
    @deserializer = hash[:unmarshaller]
    @identifier = hash[:identifier]
    @go_type = hash[:goType]
    @go_field_name = hash[:identifier].to_pascal
    @go_field_tag = "#{hash[:identifier]}_json_tag".to_camel
    @description = hash[:description] || ''
  end

  # go_struct will verify if the hash is typed as a go struct.
  def go_struct?
    return false if GO_TYPES.include?(hash[:goType])
    return false if hash[:goType].include?('scalar')

    true
  end

  def go_protofield_name
    return go_field_name unless go_struct?

    "Proto#{go_field_name}"
  end
end
