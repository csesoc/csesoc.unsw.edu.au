<template>
  <div @click.stop="hide" class="dimmed">
    <div class="vertical-container">
      <a href="/#/about">About</a>
      <a href="#">Shop</a>
      <a href="#">Events</a>
      <a href="#">CSE Resources</a>
      <a href="/#/engage">Engage</a>
      <a href="#">Sponsors</a>
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
  margin-top: 64px; /* Block out nav bar */
}

.vertical-container {
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  width: 700px;
  margin: auto;
}

.vertical-container a {
  color: white;
  padding: 2vh; /* Add some padding */
  text-decoration: none; /* Remove underline from links */
  text-align: center;
  font-size: 4vh;
  font-weight: normal;
  transition: all .2s ease-in-out;
}

.vertical-container a:hover {
  color: white;
  font-weight: bolder;
  font-size: 8vh;
}
</style>
