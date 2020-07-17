<template>
  <v-container>
    <v-list two-line>
      <template v-for="(item, index) in items">
        <v-list-item :key="item.title">
          <v-list-item-content>
            <v-list-item-title v-html="item.title"></v-list-item-title>
            <v-list-item-subtitle v-html="item.description"></v-list-item-subtitle>
          </v-list-item-content>
          <v-divider :key="index" />
        </v-list-item>
      </template>
    </v-list>
  </v-container>
</template>

<script>
export default {
  data: () => ({
    items: [
      {
        title: 'First Year Guide.',
        subtitle: 'This is a guide for students looking for guidance.'
      },
      {
        title: 'Jobs Board.',
        subtitle: 'A place where CSESOC students can look for relevant job opportunities.'
      },
      {
        title: 'Media\'s Website',
        subtitle: 'A cornucopia of articles, videos and podcasts to satisfy your desire.'
      },
      {
        title: 'Notangles.',
        subtitle: 'CSESOC\'s very own Bojangles'
      },
    ],
  }),
};
</script>

<style lang='stylus' scoped>

</style>
