docker run -d --name kong-go-plugin-geoip \
  -p "8000-8001:8000-8001" \
  -e "KONG_ADMIN_LISTEN=0.0.0.0:8001" \
  -e "KONG_PROXY_LISTEN=0.0.0.0:8000" \
  -e "KONG_DATABASE=off" \
  -e "KONG_PLUGINS=bundled,kong-go-plugin-geoip" \
  -e "KONG_PLUGINSERVER_NAMES=kong-go-plugin-geoip" \
  -e "KONG_PLUGINSERVER_KONG_GO_PLUGIN_GEOIP_QUERY_CMD=kong-go-plugin-geoip -dump" \
  xuqingfeng/kong-go-plugin-geoip