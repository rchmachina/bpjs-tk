<template>
  <div class="p-20 ml-20">>
    <label for="options">Choose an option:</label>
    <select v-model="selectedOption" id="options">
      <option disabled value="">Please select one</option>
      <option v-for="(option, index) in options" :key="index" :value="option">
        {{ option }}
      </option>
    </select>

    <p>Selected option: {{ selectedOption }}</p>

    <input v-model="searchTerm" @input="search" placeholder="Search..." />

    <table class="table w-full">
      <thead>
        <tr>
          <th class="border px-4 py-2">Sex</th>
          <th class="border px-4 py-2">No of Graduates</th>
          <th class="border px-4 py-2">Course Type</th>
          <th class="border px-4 py-2">Year</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in filteredData" :key="item.id">
          <td class="border px-4 py-2">{{ item.sex }}</td>
          <td class="border px-4 py-2">{{ item.no_of_graduates }}</td>
          <td class="border px-4 py-2">{{ item.course_type }}</td>
          <td class="border px-4 py-2">{{ item.year }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import { ref,computed} from 'vue';
import axios from 'axios';

export default {
  name: 'Soal1View',
  setup() {
    const selectedOption = ref(""); // The default value for the select box
    const options = ref([
      "Option 1",
      "Option 2",
      "Option 3",
      "Option 4",
    ]);
    
    const data = ref([]);
    const searchTerm = ref("");

    const search = () => {
      axios.get('https://data.gov.sg/api/action/datastore_search?resource_id=eb8b932c-503c-41e7-b513-114cffbe2338&limit=20', {
        params: {
          q: searchTerm.value
        }
      })
      .then(response => {
        data.value = response.data.result.records;
      })
      .catch(error => {
        console.error(error);
      });
    };

    const filteredData = computed(() => {
      if (!searchTerm.value) {
        return data.value; // Return all data if no search term
      }
      return data.value.filter(item => {
        // Modify the filtering logic as needed based on your data structure
        return item.someField.includes(searchTerm.value); // Change 'someField' to the actual field you want to filter
      });
    });

    return {
      selectedOption,
      options,
      searchTerm,
      data,
      search,
      filteredData
    };
  },
};
</script>
