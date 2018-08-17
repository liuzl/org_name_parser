go get github.com/liuzl/d/cmd/csv2dict
csv2dict -dict_dir dicts -dst company -src csv/county.txt -tag loc_county
csv2dict -dict_dir dicts -dst company -src csv/city.txt -tag loc_city
csv2dict -dict_dir dicts -dst company -src csv/province.txt -tag loc_province

csv2dict -dict_dir dicts -dst company -src csv/type.txt -tag company_type
csv2dict -dict_dir dicts -dst company -src csv/scope.txt -tag company_scope

