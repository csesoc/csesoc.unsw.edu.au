<!--
  EventGrid
  --
  This component contains a grid of event tiles containing basic information regarding each event.
  It is currently being used in the landing page.
  --
  Props:
    - events: list of event objects - containing the fields: id, name, start_time, fb_event_id, fb_cover_img
-->

<template>
  <v-container fluid>
    <!-- Catch a lack of events. -->
    <div v-if="events.length == 0">Stay tuned to our Facebook page for upcoming events!</div>
    <v-row class="justify-center">
      <v-col sm="8" md="6" lg="6" xl="4" v-for="event in events" :key="event.id">
        <v-card style = "position: relative" class="ma-1 card" height="100%"  target="_blank">
          <v-img
            :src="event.fb_cover_img"
            class="grid-img"
            height = "150px"
            gradient="to bottom, rgba(255,255,255,0.1) 0%, rgba(41,41,41,0.6) 50%, rgba(24,24,24,0.8) 100%"
          >
          </v-img>
          <v-card-text height="100%">
            <div class="name" v-text="event.name"/>
            <!-- <div class="event__place" v-if="event.place != undefined" v-text="'@ '+event.place"/> -->
            <!-- <div class="description" v-text="event.description"/> -->
          </v-card-text>
          <div class="date">
            <div class="day" v-text ="new Date(event.start_time*1000).getDate()"></div>
            <div class="month" v-text="getMonthString(event.start_time)"></div>
          </div>
          <v-card-actions>
            <a :href="'https://facebook.com/' + event.fb_event_id"><div class="link">Learn more ▶</div></a>
          </v-card-actions>

        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script type="text/javascript">
export default {
  name: 'EventGrid',
  // Must be passed from parent object
  // items have title, image url (src), and link
  props: ['events'],
  methods: {
    getMonthString(unixT) {
      const months = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sept', 'Oct', 'Nov', 'Dec'];
      const d = new Date(unixT * 1000);
      return months[d.getMonth()];
    }
    // Removed, since this is no longer used.
    /* getDateString(unix1, unix2) {
    //   const days = ["Sun", "Mon", "Tues", "Wed", "Thurs", "Fri", "Sat"];
    //   const months = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sept", "Oct", "Nov", "Dec"];

    //   if (!String.prototype.format) {
    //     String.prototype.format = function() {
    //       var args = arguments;
    //       return this.replace(/{(\d+)}/g, function(match, number) {
    //         return typeof args[number] != 'undefined'
    //           ? args[number]
    //           : match
    //         ;
    //       });
    //     };
    //   }

    //   var start = new Date(unix1 * 1000);
    //   var end = new Date(unix2 * 1000);
    //   if (unix2 == undefined) {
    //     // eg. Thurs 1 Jan 12 PM
    //     return "{0} {1} {2} at {3}{4} {5}".format(
    //       days[start.getDay()],
    //       start.getDate(),
    //       months[start.getMonth()],
    //       (start.getHours() > 12) ? start.getHours() - 12 : start.getHours(),
    //       (start.getMinutes() === 0) ? "" : ":" + (start.getMinutes() < 10 ? "0" : "") + start.getMinutes(),
    //       (start.getHours() >= 12) ? "PM" : "AM"
    //     );
    //   }
    //   // if the date is the same, we just need to display start and end times of the event
    //   else if (start.getDate() === end.getDate() && start.getMonth() === end.getMonth()) {
    //     // eg. Thurs 1 Jan 12 PM to 1:30 PM
    //     return "{0} {1} {2} at {3}{4} {5} - {6}{7} {8}".format(
    //       days[start.getDay()],
    //       start.getDate(),
    //       months[start.getMonth()],
    //       (start.getHours() > 12) ? start.getHours() - 12 : start.getHours(),
    //       (start.getMinutes() === 0) ? "" : ":" + (start.getMinutes() < 10 ? "0" : "") + start.getMinutes(),
    //       (start.getHours() >= 12) ? "PM" : "AM",
    //       (end.getHours() > 12) ? end.getHours() - 12 : end.getHours(),
    //       (end.getMinutes() === 0) ? "" : ":" + (end.getMinutes() < 10 ? "0" : "") + end.getMinutes(),
    //       (end.getHours() >= 12) ? "PM" : "AM"
    //     );
    //   }
    //   // if the dates are different
    //   // eg. Thurs 1 Jan 12 PM to Thurs 2 Jan 1:30 PM
    //   return "{0} {1} {2} at {3}{4} {5} - {6} {7} {8} at {9}{10} {11}".format(
    //     days[start.getDay()],
    //     start.getDate(),
    //     months[start.getMonth()],
    //     (start.getHours() > 12) ? start.getHours() - 12 : start.getHours(),
    //     (start.getMinutes() === 0) ? "" : ":" + (start.getMinutes() < 10 ? "0" : "") + start.getMinutes(),
    //     (start.getHours() >= 12) ? "PM" : "AM",
    //     days[end.getDay()],
    //     end.getDate(),
    //     months[end.getMonth()],
    //     (end.getHours() > 12) ? end.getHours() - 12 : end.getHours(),
    //     (end.getMinutes() === 0) ? "" : ":" + (end.getMinutes() < 10 ? "0" : "") + end.getMinutes(),
    //     (end.getHours() >= 12) ? "PM" : "AM"
    //   );
    // } */
  }
};
</script>

<style scoped>
.grid-square,
.grid-img {
  height: 28vh;
  background: rgb(8, 72, 255);
  background: linear-gradient(
    120deg,
    rgba(8, 72, 255, 1) 0%,
    rgba(0, 98, 214, 1) 50%,
    rgba(65, 115, 255, 1) 100%
  );
}

/* .description {
  color: black !important;
  font-size: 1.1em;
} */

/* .place {
  color: black !important;
  font-size: 1.1em;
} */

a {
  text-decoration: none;
}

.card {
  padding-bottom: 20px;
  min-height: 295px;
  overflow: hidden;
}

.link {
  color: rgb(54, 119, 243) !important;
  font-size: 1.2em;
  position: absolute;
  bottom: 20px;
  right: 15px;
  width: 100%;
  text-align: right;
  font-weight: bold;
  transition: color 0.3s;
}

.link:hover {
  color: rgb(97, 157, 246) !important;
}

.name {
  color: #111;
  width: 75%;
  font-size: 2em;
  font-weight: bold;
  line-height: 1.2;
}

.date {
  color: rgb(197, 197, 197);
  position: absolute;
  right: -5px;
  bottom: 45px;
  width: 100%;
  text-align: right;
}

.month {
  color: rgb(202, 202, 202);
  font-weight: bold;
  line-height: 0.2;
  font-size: 55px;
}

.day {
  font-size: 65px;
}
</style>
