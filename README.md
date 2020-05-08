### golang-vuejs

Chỉnh sửa file setting.json, vì project này dùng go-module để quản lí
```
{
    "go.useLanguageServer": true,
    "go.toolsEnvVars": {
        "GO111MODULE": "on"
    }
}
```

```
npm install js-cookie --save
npm install --save vue-resource
npm install --save vue-router
```

### Vuex state is undefined when refresh page.Fix:
```
npm install --save vuex-persistedstate
```

### multi language
```
npm install vue-i18n
```


### Fix error block cors 
```
Add header,method,origins from server

headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Accept", "User-Agent", "Origin", "auth_token", "X-CSRF-Token"})
methods := handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "PATCH"})
origins := handlers.AllowedOrigins([]string{"*"})
```






