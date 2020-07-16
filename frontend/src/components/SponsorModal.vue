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
<<<<<<< HEAD
      //Value here is the prop that is made from v-model
=======
>>>>>>> 9610311... fixed modal
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
