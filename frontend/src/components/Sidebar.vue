<template>
  <v-navigation-drawer v-model="drawer" app class="sidebar">
    <v-container fluid>
      <v-row align="center" justify="center">
        <div class="search">
          <input @input="search_text()" v-model="search.text" type="text" placeholder="Search..." />
          <v-list-item-action>
            <v-icon v-text="searchIcon"></v-icon>
          </v-list-item-action>
        </div>
      </v-row>

      <v-divider></v-divider>
    </v-container>

    <v-list>
      <div v-for="item in items" :key="item.id">
        <router-link :to="item.route" class="route">
          <v-list-item v-ripple v-if="item.children.length == 0">
            <v-list-item-icon>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-item-icon>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item>
        </router-link>

        <v-list-group v-if="item.children.length > 0" :prepend-icon="item.icon" no-action>
          <template v-slot:activator>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </template>
        </v-list-group>
      </div>
    </v-list>
  </v-navigation-drawer>
</template>

<style scoped>
.sidebar {
  font-family: "Quicksand", sans-serif;
}
</style>

<script>
export default {
  data() {
    return {
      items: [
        {
          title: 'Home',
          icon: 'mdi-console-line',
          route: '/',
          children: []
        },
        {
          title: 'About',
          icon: 'mdi-information',
          route: '/about',
          children: [
            { title: 'Who are we?', route: '/about', icon: '' },

            { title: 'FAQ', route: '/about', icon: '' },
            { title: 'Constitution', route: '/', icon: '' }
          ]
        },
        {
          title: 'Members',
          icon: 'mdi-account',
          route: '/members',
          children: [
            { title: 'Our team', route: '/', icon: '' },
            { title: 'Student Network', route: '/', icon: '' }
          ]
        },
        {
          title: 'Media',
          icon: 'mdi-camera',
          route: '/media',
          children: [
            { title: 'Videos', route: '/', icon: '' },
            { title: 'Blog', route: '/', icon: '' },
            { title: 'Podcast', route: '/', icon: '' },
            { title: 'First Year Guide', route: '/', icon: '' }
          ]
        },
        {
          title: 'Merch',
          icon: 'mdi-shopping',
          route: '/merch',
          children: []
        },
        {
          title: 'Events',
          icon: 'mdi-calendar',
          route: '/events',
          children: []
        },
        {
          title: 'Resources',
          icon: 'mdi-library-books',
          route: '/resources',
          children: []
        },
        {
          title: 'Projects',
          icon: 'mdi-laptop',
          route: '/projects',
          children: []
        },
        {
          title: 'For Sponsors',
          icon: 'mdi-city',
          route: '/sponsors',
          children: []
        },
        {
          title: 'Contact',
          icon: 'mdi-contact-mail',
          route: '/contact',
          children: []
        }
      ],
      search: { filter: null, text: '' },
      searchIcon: 'mdi-magnify'
    };
  },
  props: {
    drawer: Boolean
  },
  methods: {
    search_text() {
      console.log('typing...');
    }
  }
};
</script>

<style>
.route {
  text-decoration: none;
}
</style>
