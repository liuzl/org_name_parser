./parse -dict_dir dicts -dict_name company -cpuprofile profile.prof > /dev/null
go tool pprof ./parse profile.prof
