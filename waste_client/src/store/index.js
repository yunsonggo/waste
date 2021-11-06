import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

const types = {
  SET_GOODSALL: "SET_GOODSALL",
  SET_GOODSHOT: "SET_GOODSHOT",
  SET_GOODSMORE: "SET_GOODSMORE",
  SET_LOCATION: "SET_LOCATION",
  SET_ADDRESS: "SET_ADDRESS",
  SET_POSITION: "SET_POSITION",
};

const state = {
  goodsAll: {},
  goodsHot: {},
  goodsMore: {},
  location: {},
  address: "",
  position: {},
};

const getters = {
  goodsAll: (state) => state.goodsAll,
  goodsHot: (state) => state.goodsHot,
  goodsMore: (state) => state.goodsMore,
  location: (state) => state.location,
  address: (state) => state.address,
  position: (state) => state.position,
};

const mutations = {
  [types.SET_GOODSALL](state, goodsAll) {
    if (goodsAll) {
      state.goodsAll = goodsAll;
    } else {
      state.goodsAll = null;
    }
  },
  [types.SET_GOODSHOT](state, goodsHot) {
    if (goodsHot) {
      state.goodsHot = goodsHot;
    } else {
      state.goodsHot = null;
    }
  },
  [types.SET_GOODSMORE](state, goodsMore) {
    if (goodsMore) {
      state.goodsMore = goodsMore;
    } else {
      state.goodsMore = null;
    }
  },
  [types.SET_LOCATION](state, location) {
    if (location) {
      state.location = location;
    } else {
      state.location = null;
    }
  },
  [types.SET_ADDRESS](state, address) {
    if (address) {
      state.address = address;
    } else {
      state.address = "";
    }
  },
  [types.SET_POSITION](state, position) {
    if (position) {
      state.position = position;
    } else {
      state.position = null;
    }
  },
};
const actions = {
  setGoodsAll: ({ commit }, goodsAll) => {
    commit(types.SET_GOODSALL, goodsAll);
  },
  setGoodsHot: ({ commit }, goodsHot) => {
    commit(types.SET_GOODSHOT, goodsHot);
  },
  setGoodsMore: ({ commit }, goodsMore) => {
    commit(types.SET_GOODSMORE, goodsMore);
  },
  setLocation: ({ commit }, location) => {
    commit(types.SET_LOCATION, location);
  },
  setAddress: ({ commit }, address) => {
    commit(types.SET_ADDRESS, address);
  },
  setPosition: ({ commit }, position) => {
    commit(types.SET_POSITION, position);
  },
};
export default new Vuex.Store({
  state,
  getters,
  mutations,
  actions,
});
