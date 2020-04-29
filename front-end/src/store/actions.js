export const SetCustomer = ({ commit }, cus) => {
  commit("SetCustomer", cus);
};

export const EditCustomer = ({ commit }, cus) => {
  commit("EditCustomer", cus);
};

export const SetToken = ({ commit }, token) => {
  commit("SetToken", token);
};

export const SetLang = ({ commit }, lang) => {
  commit("SetLang", lang);
};
