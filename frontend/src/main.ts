import { createApp } from 'vue'
import App from '@/App.vue'
import router from '@/plugins/router'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
import '@/css/index.css'
import { configStore } from '@/plugins/store'

const app = createApp(App)

const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)
app.use(pinia)

app.use(router)

app.mount('#app')

const config = configStore()
config.load()
