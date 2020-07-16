<template>
  <div class="dimmed">
    <div class="vertical-menu" >
      <a @click.stop="hide" href="/#/about">About</a>
      <a @click.stop="hide" href="#">Shop</a>
      <a @click.stop="hide" href="#">Events</a>
      <a @click.stop="hide" href="#">CSE Resources</a>
      <a @click.stop="hide" href="/#/engage">Engage</a>
      <a @click.stop="hide" href="#">Sponsors</a>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      
    };
  },
  methods: {
    hide(event) {
      // Relaying status to parent
      this.$emit('shown', false)
    }
  }
};
</script>

<style>
.dimmed {
  z-index: 2;
  display: block;
  position: fixed;
  height: 100%;
  width: 100%;
  top: 0;
  left: 0;
  right: 0;
  background: rgba(0, 0, 0, 0.9);
  font-family: "Quicksand", sans-serif;
}

.vertical-menu {
  width: 500px; /* Set a width if you like */
  margin: auto;
  padding-top: 200px;
}

.vertical-menu a {
  color: #1976D2;
  display: block; /* Make the links appear below each other */
  padding: 12px; /* Add some padding */
  text-decoration: none; /* Remove underline from links */
  text-align: center;
  font-size: 35px;
  font-weight: bold;
}

.vertical-menu a:hover {
  color: white;
}
</style>
