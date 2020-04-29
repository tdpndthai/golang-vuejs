import Vue from "vue";
import vi from "../langs/vi.json";
import en from "../langs/en.json";
import VueI18n from "vue-i18n";

Vue.use(VueI18n);

export const i18n = new VueI18n({
  locale: "vi",
    fallbackLocale: "en",
    messages: {
        vi,
        en
    }
  
});
