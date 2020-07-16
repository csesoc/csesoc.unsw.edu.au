<template>
    <v-dialog
        v-model="show"
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
      value: Boolean
    },
    computed: {
      show: {
        get () {
          return this.value
        },
        set (value) {
          this.$emit('input', value)
        }
      }
    }
  }
</script>
<style scoped>
</style>
