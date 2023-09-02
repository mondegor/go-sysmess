package mrerr

var (
    FactoryInternal = NewFactory(
        ErrorIdInternal, ErrorKindInternal, "internal server error")

    FactoryInternalNilPointer = NewFactory(
        "errInternalNilPointer", ErrorKindInternal, "nil pointer")

    FactoryInternalTypeAssertion = NewFactory(
        "errInternalTypeAssertion", ErrorKindInternal, "invalid type '{{ .type }}' assertion (value: {{ .value }})")

    FactoryInternalInvalidType = NewFactory(
        "errInternalInvalidType", ErrorKindInternal, "invalid type '{{ .type1 }}', expected: '{{ .type2 }}'")

    FactoryInternalInvalidData = NewFactory(
        "errInternalInvalidData", ErrorKindInternal, "invalid data '{{ .value }}'")

    FactoryInternalParseData = NewFactory(
        "errInternalParseData", ErrorKindInternal, "data '{{ .name1 }}' parsed to {{ .name2 }} with error")

    FactoryInternalFailedToClose = NewFactory(
        "errInternalFailedToClose", ErrorKindInternal, "failed to close '{{ .name }}'")

    FactoryInternalMapValueNotFound = NewFactory(
        "errInternalMapValueNotFound", ErrorKindInternal, "'{{ .value }}' is not found in map {{ .name }}")

    FactoryDataContainer = NewFactory(
        "errDataContainer", ErrorKindInternalNotice, "data: '{{ .value }}'")
)
