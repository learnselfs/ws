import { ref } from 'vue'
const App = {
  setup(){

  }
};
const app = Vue.createApp(App);
for ([name, comp] of Object.entries(ElementPlusIconsVue)) {
  app.component(name, comp);
}
app.use(ElementPlus);
app.mount("#app");
