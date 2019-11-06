<template>
    <v-navigation-drawer v-model="drawer" app class="sidebar">

      <v-container fluid>

        <v-row align="center" justify="center">
          <div class="search">
            <input
              @input="search_text()"
              v-model="search.text"
              type="text"
              placeholder="Search..."
            />
            <v-list-item-action>
              <v-icon v-text="searchIcon"></v-icon>
            </v-list-item-action>
          </div>
        </v-row>

        <v-divider></v-divider>
      </v-container>

      <v-list>

      <!-- <router-link to="/" class="route">
      <v-list-item v-ripple>
        <v-list-item-icon>
          <v-icon>mdi-home</v-icon>
        </v-list-item-icon>
        <v-list-item-title>Home</v-list-item-title>
      </v-list-item>
      </router-link>

       <router-link to="/about" class="route">
       <v-list-item v-ripple>
         <v-list-item-icon>
           <v-icon>mdi-information</v-icon>
         </v-list-item-icon>
         <v-list-item-title>About</v-list-item-title>
       </v-list-item>
       </router-link>

       <v-list-group
         prepend-icon="mdi-account"
         no-action
         value="true"
       >
       <template v-slot:activator>
         <v-list-item-title>Users</v-list-item-title>
       </template>


         <v-list-item href="#">
           <v-list-item-title>ffffff</v-list-item-title>
         </v-list-item> -->




<!-- ---------------------------------- -->
<div v-for="item in items">
<router-link :to="item.route" class="route">
<v-list-item v-ripple v-if="item.children.length == 0">
  <v-list-item-icon>
    <v-icon>{{ item.icon }}</v-icon>
  </v-list-item-icon>
  <v-list-item-title>{{ item.title }}</v-list-item-title>
</v-list-item>
</router-link>

<v-list-group
  v-if="item.children.length > 0"
  :prepend-icon="item.icon"
  no-action
>
<template v-slot:activator>
  <v-list-item-title>{{ item.title }}</v-list-item-title>
</template>

<!-- <router-link :to="child.route" class="route"> -->
  <v-list-item v-for="child in item.children">
    <v-list-item-title>{{ child.title }}</v-list-item-title>
  </v-list-item>
<!-- </router-link> -->

</v-list-group>

</div>



         <!-- <v-list-group
           no-action
           sub-group
           value="true"
         >
         <v-list-item>
           <v-list-item-icon>
             <v-icon>mdi-home</v-icon>
           </v-list-item-icon>

           <v-list-item-title>wwww</v-list-item-title>
         </v-list-item>
           <template v-slot:activator>
             <v-list-item-content>
               <v-list-item-title>Admin</v-list-item-title>
             </v-list-item-content>
           </template>

           <v-list-item
             v-for="(admin, i) in admins"
             :key="i"
             link
           >
             <v-list-item-title v-text="admin[0]"></v-list-item-title>

           </v-list-item>
         </v-list-group> -->

         <!-- <v-list-group
           sub-group
           no-action
         >
           <template v-slot:activator>
             <v-list-item-content>
               <v-list-item-title>Actions</v-list-item-title>
             </v-list-item-content>
           </template>
           <v-list-item
             v-for="(crud, i) in cruds"
             :key="i"
             @click=""
           >
             <v-list-item-title v-text="crud[0]"></v-list-item-title>
             <v-list-item-action>

             </v-list-item-action>
           </v-list-item>
         </v-list-group> -->

       </v-list-group>
     </v-list>
    </v-navigation-drawer>
</template>

<style scoped>
@import url("https://fonts.googleapis.com/css?family=Quicksand");

.sidebar {
  font-family: "Quicksand", sans-serif;
}
</style>

<script>
export default {
  data() {
    return {
      items: [
        { title: 'Home', icon: 'mdi-console-line', route: '/', children: [] },
        {
          title: 'About',
          icon: 'mdi-information',
          route: '/about',
          children: [
            { title: 'History', route: '/', icon: '' },
            { title: 'FAQ', route: '/about', icon: '' },
            { title: 'Constitution', route: '/', icon: '' },
          ],
        },
        { title: 'Members', icon: 'mdi-account', route: '/', children: [
            { title: 'Our team', route: '/', icon: '' },
            { title: 'Student Network', route: '/', icon: '' },
        ] },
        { title: 'Media', icon: 'mdi-camera', route: '/', children: [
            { title: 'Videos', route: '/', icon: '' },
            { title: 'Blog', route: '/', icon: '' },
            { title: 'Podcast', route: '/', icon: '' },
            { title: 'First Year Guide', route: '/', icon: '' },
        ] },
        { title: 'Merch', icon: 'mdi-shopping', route: '/', children: [] },
        { title: 'Events', icon: 'mdi-calendar', route: '/about', children: [] },
        { title: 'Resources', icon: 'mdi-library-books', route: '/', children: [] },
        { title: 'Projects', icon: 'mdi-laptop', route: '/', children: [] },
        { title: 'For Sponsors', icon: 'mdi-city', route: '/', children: [] },
        { title: 'Contact', icon: 'mdi-contact-mail', route: '/', children: [] },
      ],
      search: { filter: null, text: '' },
      searchIcon: 'mdi-magnify',
    };
  },
  props: {
    drawer: Boolean,
  },
  methods: {
    search_text() {
      console.log('typing...');
    },
  },
};
</script>

<style>
    .route {
        text-decoration: none;
    }
</style>
