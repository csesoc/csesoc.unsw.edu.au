<!--
  HeaderTitle
  --
  This component simply formats the title. Ensure this component
  is used for all titles to keep consistency across the site.
  --
  Props:
    - title: string
-->

<template>
  <div class="header-title">
    <h1 v-bind:style="{padding: padding}">{{ title }}</h1>
    <h4 v-if="subtitle !== false">{{ subtitle }}</h4>
  </div>
</template>

<script>
export default {
  name: 'HeaderTitle',
  props: {
    title: {
      type: String,
      required: true,
    },
    subtitle: {
      default: false,
    },
    padding: {
      required: false,
    }
  }
};
</script>

<style scoped>
@media only screen and (max-width: 450px) {
  .header-title > h1 {
    font-size: 2.5em;
  }
}
</style>
