<template>
  <v-container>
    <v-row align="center" justify="center">
      <v-col>
        <v-row justify="center">
          <v-container class="preview-container">
            <router-link to="/">
              <v-img :src="`data:image/png;base64,${preview}`"/>
            </router-link>
          </v-container>
        </v-row>
      </v-col>
      <v-col>
        <v-list two-line>
          <template v-for="(item, index) in items">
            <v-list-item @mouseover="preview = item.image" :key="item.title">
              <v-list-item-content>
                <v-list-item-title v-html="item.title"></v-list-item-title>
                <v-list-item-subtitle v-html="item.description"></v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
            <v-divider :key="index" />
          </template>
        </v-list>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'Preview',
  props: ['items'],
  data: () => ({
    preview: '@/assets/csesocgreyblue.png',
    items: [
      {
        title: 'First Year Guide.',
        description: 'This is a guide for students looking for guidance.'
      },
      {
        title: 'Jobs Board.',
        description: 'A place where CSESOC students can look for relevant job opportunities.'
      },
      {
        title: 'Media\'s Website',
        description: 'A cornucopia of articles, videos and podcasts to satisfy your desire.'
      },
      {
        title: 'Notangles.',
        description: 'CSESOC\'s very own Bojangles'
      },
    ],
  }),
};
</script>

<style scoped>
  .preview-container {
    max-width: 600px;
  }
</style>
