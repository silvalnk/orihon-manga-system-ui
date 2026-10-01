import { createApp, h, type DefineComponent } from 'vue'
import { createInertiaApp } from '@inertiajs/vue3'
import { createPinia } from 'pinia'
import './styles/app.css'

showDesktopPointer()

const pinia = createPinia()

createInertiaApp({
  resolve: (name) => {
    const pages = import.meta.glob<DefineComponent>('./presentation/pages/*.vue', { eager: true })
    const page = pages[`./presentation/pages/${name}.vue`]
    if (!page) throw new Error(`missing page ${name}`)
    return page
  },
  setup({ el, App, props, plugin }) {
    createApp({ render: () => h(App, props) })
      .use(plugin)
      .use(pinia)
      .mount(el)
  },
})

function showDesktopPointer() {
  if (!navigator.userAgent.includes('wails.io')) return
  document.documentElement.classList.add('orihon-pointer')
  const pointer = document.createElement('div')
  pointer.setAttribute('aria-hidden', 'true')
  pointer.style.position = 'fixed'
  pointer.style.left = '0'
  pointer.style.top = '0'
  pointer.style.zIndex = '10000'
  pointer.style.pointerEvents = 'none'
  pointer.style.opacity = '0'
  pointer.innerHTML =
    '<svg width="18" height="18" viewBox="0 0 18 18" xmlns="http://www.w3.org/2000/svg"><path d="M1 1 L1 14 L4.6 10.4 L7.4 16 L10.2 14.8 L7.3 9.4 L12.2 9.2 Z" fill="#1c1915" stroke="#f4efe4" stroke-width="1"/></svg>'
  document.addEventListener('mousemove', (event) => {
    pointer.style.opacity = '1'
    pointer.style.transform = `translate(${event.clientX}px, ${event.clientY}px)`
  })
  document.addEventListener('mouseleave', () => {
    pointer.style.opacity = '0'
  })
  document.body.appendChild(pointer)
}
