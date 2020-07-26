<template>
  <v-container class="nav__grid" fluid>
    <!-- Catch a lack of events. -->
    <div v-if="events.length == 0">Stay tuned to our Facebook page for upcoming events!</div>
    <v-row class="justify-center">
      <v-col sm="12" md="6" lg="6" xl="4" v-for="event in events" :key="event.id">
        <v-card class="ma-1 grid__square" height="100%" :href="'https://facebook.com/'+ event.fb_event_id" target="_blank">
          <v-img
            :src="event.fb_cover_img"
            class="grid__img"
            gradient="to bottom, rgba(255,255,255,0.1) 0%, rgba(41,41,41,0.6) 50%, rgba(24,24,24,0.8) 100%"
          >
            <v-card-title
              class="justify-center fill-height align-center white--text font-weight-medium text-truncate"
              v-text="event.name"
            ></v-card-title>
          </v-img>
          <v-card-text>
            <div class="event__date" v-text="getDateString(event.start_time, event.end_time)"/>
            <div class="event__place" v-if="event.place != undefined" v-text="'@ '+event.place"/>
            <div class="event__description" v-text="event.description"/>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>


<style scoped>
.grid_square,
.grid__img {
  height: 28vh;
  background: rgb(8, 72, 255);
  background: linear-gradient(
    120deg,
    rgba(8, 72, 255, 1) 0%,
    rgba(0, 98, 214, 1) 50%,
    rgba(65, 115, 255, 1) 100%
  );
}
.event__description {
  color: black !important;
  font-size: 1.1em;
}
.event__place {
  color: black !important;
  font-size: 1.1em;
}
.event__date {
  color: rgb(54, 119, 243) !important;
  font-size: 1.2em;
  font-weight: bold;
}
</style>

<script type="text/javascript">
export default {
  name: 'EventGrid',
  // Must be passed from parent object
  // items have title, image url (src), and link
  props: ['events'],
  methods: {
    getDateString(unix1, unix2) {
      const days = ["Sun", "Mon", "Tues", "Wed", "Thurs", "Fri", "Sat"];
      const months = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sept", "Oct", "Nov", "Dec"];

      if (!String.prototype.format) {
        String.prototype.format = function() {
          var args = arguments;
          return this.replace(/{(\d+)}/g, function(match, number) { 
            return typeof args[number] != 'undefined'
              ? args[number]
              : match
            ;
          });
        };
      }
     
      var start = new Date(unix1 * 1000);
      var end = new Date(unix2 * 1000);
      if (unix2 == undefined) {
        // eg. Thurs 1 Jan 12 PM
        return "{0} {1} {2} at {3}{4} {5}".format(
          days[start.getDay()], 
          start.getDate(),
          months[start.getMonth()],
          (start.getHours() > 12) ? start.getHours() - 12 : start.getHours(),
          (start.getMinutes() === 0) ? "" : ":" + (start.getMinutes() < 10 ? "0" : "") + start.getMinutes(),
          (start.getHours() >= 12) ? "PM" : "AM"
        );
      }
      // if the date is the same, we just need to display start and end times of the event
      else if (start.getDate() === end.getDate() && start.getMonth() === end.getMonth()) {
        // eg. Thurs 1 Jan 12 PM to 1:30 PM
        return "{0} {1} {2} at {3}{4} {5} - {6}{7} {8}".format(
          days[start.getDay()], 
          start.getDate(),
          months[start.getMonth()],
          (start.getHours() > 12) ? start.getHours() - 12 : start.getHours(),
          (start.getMinutes() === 0) ? "" : ":" + (start.getMinutes() < 10 ? "0" : "") + start.getMinutes(),
          (start.getHours() >= 12) ? "PM" : "AM",
          (end.getHours() > 12) ? end.getHours() - 12 : end.getHours(),
          (end.getMinutes() === 0) ? "" : ":" + (end.getMinutes() < 10 ? "0" : "") + end.getMinutes(),
          (end.getHours() >= 12) ? "PM" : "AM"
        );
      } 
      // if the dates are different
      // eg. Thurs 1 Jan 12 PM to Thurs 2 Jan 1:30 PM
      return "{0} {1} {2} at {3}{4} {5} - {6} {7} {8} at {9}{10} {11}".format(
        days[start.getDay()], 
        start.getDate(),
        months[start.getMonth()],
        (start.getHours() > 12) ? start.getHours() - 12 : start.getHours(),
        (start.getMinutes() === 0) ? "" : ":" + (start.getMinutes() < 10 ? "0" : "") + start.getMinutes(),
        (start.getHours() >= 12) ? "PM" : "AM",
        days[end.getDay()], 
        end.getDate(),
        months[end.getMonth()],
        (end.getHours() > 12) ? end.getHours() - 12 : end.getHours(),
        (end.getMinutes() === 0) ? "" : ":" + (end.getMinutes() < 10 ? "0" : "") + end.getMinutes(),
        (end.getHours() >= 12) ? "PM" : "AM"
      );
    }
  }
};
</script>
