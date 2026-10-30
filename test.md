go test ./otp -v

go test ./aleatorio -run TestBytesAleatorios -v

go test ./otp -run TestCifrarYDescifrar -v
go test ./otp -run TestCifrarYDescifrar -v -count=1