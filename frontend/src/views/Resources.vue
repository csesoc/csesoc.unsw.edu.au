<template>
  <v-container>
    <h1>Heading 1</h1>
    <h2>Heading 2</h2>
    <h3>Heading 3</h3>
    <h4>Heading 4</h4>
    <p>Paragraph</p>
    <!-- <v-row>
      <v-col class="pa-4" cols='6'>
        Some text
      </v-col>
      <v-col class="pa-4" cols='6'>
        Some other text
      </v-col>
    </v-row> -->
    <div class="flex-grid">
      <div class="col col--md-12">100%</div>
      <div class="col col--md-6">50%</div>
      <div class="col col--md-6">50%</div>
      <div class="col col--md-4">33%</div>
      <div class="col col--md-4">33%</div>
      <div class="col col--md-4">33%</div>
      <div class="col col--md-3">25%</div>
      <div class="col col--md-3">25%</div>
      <div class="col col--md-3">25%</div>
      <div class="col col--md-3">25%</div>
      <div class="col col--md-4">33%</div>
      <div class="col col--md-8">66%</div>
      <div class="col col--md-8">66%</div>
      <div class="col col--md-4">33%</div>
      <div class="col col--md-3">25%</div>
      <div class="col col--md-6">50%</div>
      <div class="col col--md-3">25%</div>
    </div>
  </v-container>
</template>

<script>
</script>

<style lang="scss" scoped>
.col {
  background-color: $primary-color;
}
</style>
