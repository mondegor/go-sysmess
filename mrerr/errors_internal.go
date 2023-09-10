package mrerr

var (
    ErrFactoryInternal = NewFactory(
        ErrorIdInternal, ErrorKindInternal, "internal server error")

    ErrFactoryInternalNilPointer = NewFactory(
        "errInternalNilPointer", ErrorKindInternal, "nil pointer")

    ErrFactoryInternalTypeAssertion = NewFactory(
        "errInternalTypeAssertion", ErrorKindInternal, "invalid type '{{ .type }}' assertion (value: {{ .value }})")

    ErrFactoryInternalInvalidType = NewFactory(
        "errInternalInvalidType", ErrorKindInternal, "invalid type '{{ .type1 }}', expected: '{{ .type2 }}'")

    ErrFactoryInternalInvalidData = NewFactory(
        "errInternalInvalidData", ErrorKindInternal, "invalid data '{{ .value }}'")

    ErrFactoryInternalParseData = NewFactory(
        "errInternalParseData", ErrorKindInternal, "data '{{ .name1 }}' parsed to {{ .name2 }} with error")

    ErrFactoryInternalFailedToClose = NewFactory(
        "errInternalFailedToClose", ErrorKindInternal, "failed to close '{{ .name }}'")

    ErrFactoryInternalMapValueNotFound = NewFactory(
        "errInternalMapValueNotFound", ErrorKindInternal, "'{{ .value }}' is not found in map {{ .name }}")

    ErrFactoryInternalNoticeDataContainer = NewFactory(
        "errInternalNoticeDataContainer", ErrorKindInternalNotice, "data: '{{ .value }}'")
)
