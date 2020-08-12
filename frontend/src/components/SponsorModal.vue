<!--
  SponsorModal
  --
  This component is a modal that pops up when a sponsor link is clicked.
  It contains information related to the selected sponsor.
  --
  Props:
    - title: string
    - body: string
    - value: boolean
-->

<template>
    <v-dialog v-model="show" max-width="500">
      <v-card>
        <v-card-title class="headline">{{title}}</v-card-title>
        <v-card-text>
            {{body}}
        </v-card-text>
      </v-card>
    </v-dialog>
</template>

<script>
export default {
  props: {
    title: String,
    body: String,
    // Value here is the prop that is made from v-model
    value: Boolean
  },
  computed: {
    show: {
      get() {
        return this.value;
      },
      set(value) {
        this.$emit('input', value);
      }
    }
  }
};
</script>

<style scoped>
</style>
