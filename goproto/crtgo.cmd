Get-ChildItem roomMessage.proto |Resolve-Path -Relative | %{protoc $_ --go_out=.}