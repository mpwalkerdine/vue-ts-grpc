import Vue from 'vue';
import Vuex from 'vuex';

import { RootState } from './types';

import {
  ListExamplesRequest,
  CreateExampleRequest,
  Example as ProtoExample,
  DeleteExampleRequest,
} from '@/_proto/example_pb';

import { ExampleServiceClient } from '@/_proto/example_pb_service';

// TODO config this
const client = new ExampleServiceClient('http://localhost:8081');

Vue.use(Vuex);

// TODO break store up into modules
export default new Vuex.Store<RootState>({
  state: {
    examples: [],
  },
  mutations: {
    SET_EXAMPLES(state, examples) {
      Vue.set(state, 'examples', examples);
    },
    ADD_EXAMPLE(state, example) {
      state.examples.push(example);
    },
    DELETE_EXAMPLE(state, name) {
      Vue.set(state, 'examples', state.examples.filter((eg) => eg.name !== name ));
    },
  },
  getters: {
    getExamples(state) {
      return state.examples;
    },
  },
  actions: {
    fetchExamples(context) {
      const req = new ListExamplesRequest();
      client.listExamples(req, (err, res) => {
        if (err) {
          // tslint:disable-next-line:no-console TODO error handling
          console.error(err);
          return;
        }
        const mapped = res!.getExamplesList().map((eg) => {
          return eg.toObject();
        });
        context.commit('SET_EXAMPLES', mapped);
      });
    },
    addExample(context) {
      const name = Math.random().toString(36).substring(2, 15) + Math.random().toString(36).substring(2, 15);
      const newItem = new ProtoExample();
      newItem.setName(name);
      newItem.setText('Example Text for ' + name);

      const req = new CreateExampleRequest();
      req.setParent('examples');
      req.setExample(newItem);

      client.createExample(req, (err, res) => {
        if (err) {
          // tslint:disable-next-line:no-console TODO error handling
          console.error(err);
          return;
        }
        if (!res) {
          // tslint:disable-next-line:no-console TODO error handling
          console.error('No response received from server');
          return;
        }
        context.commit('ADD_EXAMPLE', res.toObject());
      });
    },
    deleteExample(context, name) {
      const req = new DeleteExampleRequest();
      req.setName(name);
      client.deleteExample(req, (err, res) => {
        if (err) {
          // tslint:disable-next-line:no-console TODO error handling
          console.error(err);
          return;
        }
        context.commit('DELETE_EXAMPLE', name);
      });
    },
  },
});
