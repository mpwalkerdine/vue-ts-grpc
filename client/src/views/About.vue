<template>
  <div class="about">
    <h1>This is an about page</h1>
    <h2>It shows data loaded from a gRPC server.</h2>
    <button @click="addExample">Add Example</button>
    <examples-list :examples=getExamples></examples-list>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch } from 'vue-property-decorator';
import { Action, State, Getter } from 'vuex-class';
import { Example } from '@/types';
import ExamplesList from '@/components/ExamplesList.vue'; // @ is an alias to /src

@Component({
  components: {
    ExamplesList,
  },
})
export default class About extends Vue {
  // TODO types
  @Action protected fetchExamples: any;
  @Action protected addExample: any;
  @Getter protected getExamples: any;

  protected mounted() {
    this.fetchExamples();
  }
}
</script>
