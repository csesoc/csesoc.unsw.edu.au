<template>
    <v-dialog
        v-model="vDialog"
        max-width="500"
        >
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
        dialog: Boolean
    },
    computed: {
      vDialog: function() {
        return this.dialog;
      }
    }
  }
</script>

<style scoped>
</style>
