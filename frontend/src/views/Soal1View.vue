


<template>
  <div class="p-20 ml-20 overflow-x-auto">
    <input class="input w-full max-w-xs mb-10" type="text" v-model="searchTerm" @input="search" placeholder="Enter your query" />
    <table class="table w-full">
      <thead>
        <tr>
          <th class="border px-4 py-2">sex</th>
          <th class="border px-4 py-2">no of graduates</th>
          <th class="border px-4 py-2">course type</th>
          <th class="border px-4 py-2">year</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in data" :key="item.id">
          <td class="border px-4 py-2">{{ item.sex }}</td>
          <td class="border px-4 py-2">{{ item.no_of_graduates }}</td>
          <td class="border px-4 py-2">{{ item.type_of_course }}</td>
          <td class="border px-4 py-2">{{ item.year }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'DataView',
  data() {
    return {
      data: [],
      searchTerm: '',
      //searchResults: []
    };
  },
  mounted() {
    this.search();
  },
  methods: {

    search() {

      axios.get('https://data.gov.sg/api/action/datastore_search?resource_id=eb8b932c-503c-41e7-b513-114cffbe2338&limit=20', {
        params: {

          q: this.searchTerm
        }
      })
        .then(response => {
          this.data =  response.data.result.records;
        })
        .catch(error => {
          console.error(error);
        });
    }

  },
};
</script>

<style scoped>
/* No additional styles needed for DaisyUI table */
</style>
