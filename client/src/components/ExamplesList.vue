<template>
    <div class="examples-list">
        <h2>Examples</h2>
        <ul>
            <li v-for="example in examples" :key="example.name">
                <span class="example-text">{{ example.text }}</span>
                <button
                    class="delete-example"
                    @click="deleteExample(example.name)">delete</button>
            </li>
        </ul>
    </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { Action, State, Getter } from 'vuex-class';
import { Example } from '@/types';

@Component
export default class ExamplesList extends Vue {
  @Prop({required: true}) protected examples!: any;
  @Action protected deleteExample: any; // TODO figure out actual type here...
}
</script>

<style lang="stylus" scoped>
.examples-list
    .delete-example
        font-weight bold
        margin-left 10px
        cursor pointer
        border solid 1px #111
        padding 0px 5px
        vertical-align middle
</style>
