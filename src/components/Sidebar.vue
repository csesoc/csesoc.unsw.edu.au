<template>
  <div id="sidebar">
    <v-navigation-drawer v-model="drawer" app>
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
      <v-list dense>
        <v-list-item
          v-for="(item, i) in items"
          :key="i"
          @click.stop="drawer = !drawer"
        >
          <router-link :to="item.link">
            <v-list-item-action>
              <v-icon v-text="item.icon"></v-icon>
            </v-list-item-action>
          </router-link>
          <v-list-item-content>
            <v-list-item-title v-text="item.title"></v-list-item-title>
          </v-list-item-content>
          <v-list-item-action v-if="item.children.length != 0">
            <v-btn icon>
              <v-icon
                color="grey lighten-1"
                @click.stop="toggleDropdown($event, item.title)"
                >mdi-chevron-down</v-icon
              >
              <v-list
                dense
                class="py-0 pl-1"
                v-if="dropDowns[item.title] == false"
              >
                <v-list-group v-for="(childItem, i) in item.children" :key="i">
                  <v-list-item-content>
                    <router-link to="/">
                      <v-list-item-title
                        v-text="childItem.title"
                      ></v-list-item-title>
                    </router-link>
                  </v-list-item-content>
                </v-list-group>
              </v-list>
            </v-btn>
          </v-list-item-action>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
  </div>
</template>
<script>
export default {
  data() {
    return {
      items: [
        {
          title: 'Home', icon: 'mdi-home', children: [], link: '/',
        },
        {
          title: 'About',
          icon: 'mdi-information',
          children: [
            { title: 'History', icon: '' },
            { title: 'FAQ', icon: '' },
            { title: 'Constitution', icon: '' },
          ],
          link: '/about',
        },
        {
          title: 'Members', icon: 'mdi-account', children: [], link: '/members',
        },
        {
          title: 'Media', icon: 'mdi-camera', children: [], link: '/media',
        },
        {
          title: 'Merch', icon: 'mdi-shopping', children: [], link: '/merch',
        },
        {
          title: 'Events', icon: 'mdi-calendar', children: [], link: '/events',
        },
        {
          title: 'Resources', icon: 'mdi-library-books', children: [], link: '/resources',
        },
        {
          title: 'Projects', icon: 'mdi-laptop', children: [], link: '/projects',
        },
        {
          title: 'Contact', icon: 'mdi-contact-mail', children: [], link: '/contact',
        },
      ],
      search: { filter: null, text: '' },
      searchIcon: 'mdi-magnify',
      dropDowns: {},
    };
  },
  props: {
    drawer: Boolean,
  },
  methods: {
    search_text() {
      console.log('typing...');
    },
    toggleDropdown: (event, dropdownName) => {
      event.stopPropagation();
      this.dropDowns[dropdownName] = !this.dropDowns[dropdownName];
    },
  },
  beforeMount() {
    for (let i = 0; i < this.items.length; i += 1) {
      console.log(this.items[i]);
      if (this.items[i].children.length !== 0) {
        this.dropDowns[this.items[i].title] = false;
      }
    }
  },
};
</script>

    <style>
#side-nav {
  height: 100%; /* 100% Full-height */
  width: 0; /* 0 width - change this with JavaScript */
  position: fixed; /* Stay in place */
  z-index: 1; /* Stay on top */
  top: 0; /* Stay at the top */
  left: 0;
  overflow-x: hidden; /* Disable horizontal scroll */
  padding-top: 60px; /* Place content 60px from the top */
  transition: 0.3s; /* 0.5 second transition effect to slide in the sidenav */
  background-color: grey;
}
#side-nav-btn {
  position: absolute;
  z-index: 1000;
}
ul {
  padding: 0;
}
li {
  list-style-type: none;
  background-color: lightgrey;
  width: 100%;
  margin: 0;
  padding: 0.75em 0 0.75em 0;
}
.collapsable p {
  display: inline;
}
.expand {
  float: right;
}
.collapse {
  display: none;
}
svg {
  cursor: pointer;
}
</style>
