<template>
  <div @click.stop="hide" class="dimmed no-scroll">
    <div class="vertical-container">
      <RouterLink to="/about" class="link">
        <label>About</label>
      </RouterLink>
      <RouterLink to="/" class="link">
        <label>Shop</label>
      </RouterLink>
      <RouterLink to="/" class="link">
        <label>Events</label>
      </RouterLink>
      <RouterLink to="/" class="link">
        <label>CSE Resources</label>
      </RouterLink>
      <RouterLink to="/engage" class="link">
        <label>Engage</label>
      </RouterLink>
      <RouterLink to="/" class="link">
        <label>Sponsors</label>
      </RouterLink>
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
  z-index: 20;
  display: block;
  position: fixed;
  height: 100%;
  width: 100%;
  top: 0;
  left: 0;
  right: 0;
  background: rgba(0, 0, 0, 1);
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

.vertical-container .link {
  text-decoration: none; /* Remove underline from links */
  text-align: center;
  color: white;
  padding: 2vh; /* Add some padding */
  cursor: pointer; /* To simulate a link */
  font-size: 4vh;
  font-weight: normal;
  transition: all .2s ease-in-out;
}

.vertical-container .link:hover {
  color: white;
  font-weight: bolder;
  font-size: 8vh;
}

label {
  cursor: pointer; /* To simulate a link */
}
</style>
