<template>
  <!-- Start: Key -->
  <div class="splash-page">
    <div
      class="key off"
      :class="{
        off: !glow,
        on: glow,
      }"
    >
      <svg viewBox="0 0 180 150">
        <path
          class="data-circle inner"
          d="M119.5,76.3c0-0.3,0.2-0.5,0.5-0.5l3.7-0.1l-0.7-0.6c-0.1-0.1-0.1-0.2-0.1-0.3c0-0.2,0.2-0.4,0.4-0.4c0.1,0,0.2,0,0.3,0.1 l1.2,1.2c0.1,0.1,0.2,0.2,0.2,0.4c0,0.3-0.2,0.6-0.5,0.6l-4.4,0.1C119.7,76.8,119.5,76.5,119.5,76.3z M118.9,81.3 c0.3-1.5,1.8-1.9,3.1-1.6c1.3,0.2,2.6,1.1,2.3,2.6c-0.3,1.5-1.8,1.9-3.1,1.6C119.9,83.7,118.7,82.8,118.9,81.3z M123.5,82.2 c0.2-0.8-0.7-1.3-1.7-1.5s-1.9,0-2.1,0.8c-0.2,0.9,0.7,1.3,1.7,1.5C122.4,83.2,123.4,83,123.5,82.2z M117.3,87.2 c0.1-0.2,0.4-0.4,0.6-0.3l3.4,1.3l-0.4-0.8c0-0.1,0-0.2,0-0.3c0.1-0.2,0.3-0.3,0.5-0.2c0.1,0,0.2,0.1,0.2,0.2l0.7,1.5 c0.1,0.1,0.1,0.3,0,0.4c-0.1,0.3-0.4,0.4-0.7,0.3l-4.1-1.6C117.3,87.7,117.2,87.4,117.3,87.2z M115,91.7c0.8-1.3,2.3-1.1,3.5-0.4 c1.1,0.7,2,2,1.2,3.3c-0.8,1.3-2.3,1.1-3.5,0.4C115,94.3,114.2,93,115,91.7z M118.9,94.1c0.5-0.7-0.1-1.5-1-2s-1.8-0.7-2.2,0 c-0.5,0.7,0.1,1.5,1,2S118.5,94.9,118.9,94.1z M111.3,96.5c0.2-0.2,0.5-0.2,0.7,0l2.7,2.5l-0.1-0.9c0-0.1,0-0.2,0.1-0.3 c0.1-0.2,0.4-0.2,0.6,0c0.1,0.1,0.1,0.2,0.1,0.3l0.1,1.7c0,0.1,0,0.3-0.2,0.4c-0.2,0.2-0.5,0.3-0.8,0l-3.3-3 C111.1,97,111.1,96.7,111.3,96.5z M107.5,99.9c1.2-0.9,2.6-0.2,3.4,0.9c0.8,1.1,1.1,2.6,0,3.5c-1.2,0.9-2.6,0.2-3.4-0.9 C106.6,102.3,106.3,100.8,107.5,99.9z M110.3,103.6c0.7-0.5,0.4-1.4-0.2-2.2c-0.6-0.8-1.4-1.3-2.1-0.8c-0.7,0.5-0.4,1.4,0.2,2.2 C108.8,103.6,109.6,104.1,110.3,103.6z M102.3,103.1c0.2-0.1,0.5,0,0.6,0.2l1.6,3.3l0.3-0.9c0-0.1,0.1-0.2,0.2-0.2 c0.2-0.1,0.4,0,0.5,0.2c0,0.1,0.1,0.2,0,0.3l-0.5,1.6c0,0.1-0.1,0.2-0.3,0.3c-0.3,0.1-0.6,0-0.7-0.2l-2-4 C102,103.5,102.1,103.2,102.3,103.1z M97.9,104.7c0.2-0.1,0.5,0.1,0.6,0.3l1.1,3.5l0.4-0.8c0-0.1,0.1-0.2,0.2-0.2 c0.2-0.1,0.4,0,0.5,0.3c0,0.1,0,0.2,0,0.3l-0.8,1.5c-0.1,0.1-0.2,0.2-0.3,0.3c-0.3,0.1-0.6-0.1-0.7-0.3l-1.3-4.2 C97.5,105.1,97.7,104.8,97.9,104.7z M93.4,105.7c0.3,0,0.5,0.1,0.5,0.4l0.5,3.6L95,109c0.1-0.1,0.2-0.1,0.2-0.1 c0.2,0,0.4,0.1,0.5,0.3c0,0.1,0,0.2-0.1,0.3l-1,1.4c-0.1,0.1-0.2,0.2-0.4,0.2c-0.3,0-0.6-0.1-0.6-0.5l-0.6-4.4 C92.9,106,93.1,105.7,93.4,105.7z M88.3,105.9c1.5,0.1,2.1,1.5,2,2.8c-0.1,1.4-0.8,2.7-2.2,2.7s-2.1-1.5-2-2.8 C86.1,107.2,86.8,105.8,88.3,105.9z M88.1,110.5c0.9,0,1.2-0.8,1.2-1.9c0-1-0.2-1.9-1.1-2c-0.9,0-1.2,0.8-1.2,1.9 C87,109.6,87.2,110.5,88.1,110.5z M81.9,104.9c1.4,0.4,1.7,1.9,1.4,3.2c-0.3,1.3-1.3,2.5-2.8,2.1c-1.4-0.4-1.7-1.9-1.4-3.2 C79.5,105.8,80.4,104.5,81.9,104.9z M80.7,109.4c0.8,0.2,1.4-0.6,1.6-1.5s0.2-1.9-0.6-2.1c-0.8-0.2-1.4,0.6-1.6,1.6 C79.8,108.3,79.9,109.2,80.7,109.4z M76.2,102.9c0.2,0.1,0.3,0.4,0.2,0.6l-1.6,3.3l0.9-0.3c0.1,0,0.2,0,0.3,0 c0.2,0.1,0.3,0.3,0.2,0.5c0,0.1-0.1,0.2-0.2,0.2l-1.6,0.6c-0.1,0.1-0.3,0-0.4,0c-0.3-0.1-0.4-0.4-0.3-0.7l1.9-4 C75.7,102.8,76,102.7,76.2,102.9z M72.2,100.5c0.2,0.1,0.3,0.4,0.1,0.7l-2.1,3l0.9-0.2c0.1,0,0.2,0,0.3,0.1c0.2,0.1,0.2,0.4,0.1,0.6 c-0.1,0.1-0.2,0.1-0.3,0.2l-1.7,0.3c-0.1,0-0.3,0-0.4-0.1c-0.2-0.2-0.3-0.5-0.2-0.8l2.5-3.6C71.7,100.4,72,100.3,72.2,100.5z M68.3,97.1c1.1,1.1,0.5,2.5-0.4,3.5c-1,1-2.4,1.5-3.5,0.4s-0.5-2.5,0.4-3.5C65.8,96.6,67.3,96.1,68.3,97.1z M65,100.4 c0.6,0.6,1.5,0.2,2.2-0.5c0.7-0.7,1.1-1.6,0.5-2.2c-0.6-0.6-1.5-0.2-2.2,0.5C64.8,99,64.4,99.8,65,100.4z M64.3,92.1 c0.8,1.3,0,2.6-1.2,3.3c-1.1,0.7-2.7,0.9-3.5-0.3c-0.8-1.3,0-2.6,1.2-3.3C61.9,91,63.5,90.8,64.3,92.1z M60.4,94.6 c0.5,0.7,1.4,0.5,2.2,0c0.9-0.5,1.4-1.3,1-2c-0.5-0.7-1.4-0.5-2.2,0C60.5,93.1,59.9,93.9,60.4,94.6z M61.5,86.7 c0.1,0.2,0,0.5-0.3,0.6l-3.4,1.4l0.9,0.3c0.1,0,0.2,0.1,0.2,0.2c0.1,0.2,0,0.4-0.2,0.5c-0.1,0-0.2,0-0.3,0l-1.6-0.7 c-0.1,0-0.2-0.2-0.3-0.3c-0.1-0.3,0-0.6,0.3-0.7l4.1-1.6C61.1,86.3,61.4,86.4,61.5,86.7z M60.2,81.8c0.3,1.5-1,2.4-2.3,2.6 c-1.3,0.3-2.8-0.1-3.1-1.6c-0.3-1.5,1-2.4,2.3-2.6C58.4,80,59.9,80.3,60.2,81.8z M55.6,82.7c0.2,0.8,1.1,1,2.1,0.8 c1-0.2,1.8-0.7,1.6-1.5c-0.2-0.8-1.1-1-2.1-0.8C56.2,81.4,55.4,81.9,55.6,82.7z M59.5,75.8c0,0.3-0.2,0.5-0.5,0.5l-3.7,0.1l0.7,0.6 c0.1,0.1,0.1,0.2,0.1,0.3c0,0.2-0.2,0.4-0.4,0.4c-0.1,0-0.2,0-0.3-0.1l-1.2-1.2c-0.1-0.1-0.2-0.2-0.2-0.4c0-0.3,0.2-0.6,0.5-0.6 l4.4-0.1C59.3,75.3,59.5,75.5,59.5,75.8z M59.9,71.1c0,0.3-0.3,0.4-0.5,0.4L55.8,71l0.6,0.7c0.1,0.1,0.1,0.2,0.1,0.3 c0,0.2-0.2,0.4-0.5,0.3c-0.1,0-0.2-0.1-0.3-0.2l-1-1.4c-0.1-0.1-0.1-0.3-0.1-0.4c0-0.3,0.3-0.5,0.6-0.5l4.4,0.6 C59.8,70.6,60,70.8,59.9,71.1z M61.1,66.5c-0.1,0.2-0.3,0.4-0.6,0.3L57,65.8l0.4,0.8c0.1,0.1,0.1,0.2,0,0.3 c-0.1,0.2-0.3,0.3-0.5,0.3c-0.1,0-0.2-0.1-0.2-0.2l-0.8-1.5c-0.1-0.1-0.1-0.3,0-0.4c0.1-0.3,0.4-0.5,0.7-0.4l4.2,1.3 C61,66,61.1,66.3,61.1,66.5z M63.1,61.9c-0.7,1.3-2.3,1.2-3.4,0.6c-1.2-0.6-2.1-1.9-1.4-3.2c0.7-1.3,2.2-1.2,3.4-0.6 C62.9,59.3,63.8,60.6,63.1,61.9z M59,59.7c-0.4,0.8,0.2,1.4,1.1,1.9c0.9,0.5,1.8,0.6,2.2-0.1c0.4-0.8-0.2-1.5-1.1-1.9 C60.4,59.1,59.4,58.9,59,59.7z M66.5,56.9c-0.2,0.2-0.5,0.2-0.7,0.1l-2.9-2.3l0.1,0.9c0,0.1,0,0.2-0.1,0.3c-0.1,0.2-0.4,0.2-0.6,0.1 c-0.1-0.1-0.1-0.2-0.1-0.3l-0.2-1.7c0-0.1,0-0.3,0.1-0.4c0.2-0.2,0.5-0.3,0.8-0.1l3.4,2.8C66.6,56.3,66.6,56.6,66.5,56.9z M69.8,53.5c-0.2,0.2-0.5,0.2-0.7,0l-2.5-2.7l0,0.9c0,0.1,0,0.2-0.1,0.3c-0.2,0.1-0.4,0.1-0.6,0c-0.1-0.1-0.1-0.2-0.1-0.3l0.1-1.7 c0-0.1,0.1-0.3,0.2-0.4c0.2-0.2,0.6-0.2,0.8,0l3,3.3C70,53,70,53.3,69.8,53.5z M73.9,50.5c-1.3,0.8-2.6-0.1-3.3-1.2 c-0.7-1.2-0.9-2.7,0.4-3.5c1.3-0.8,2.6,0.1,3.3,1.2C75,48.2,75.2,49.7,73.9,50.5z M71.4,46.5c-0.7,0.4-0.5,1.4,0,2.2 c0.5,0.9,1.3,1.5,2,1c0.7-0.5,0.5-1.4,0-2.2C72.9,46.7,72.2,46.1,71.4,46.5z M79.7,47.7c-1.4,0.5-2.5-0.6-2.9-1.9 c-0.4-1.3-0.3-2.8,1.1-3.3c1.4-0.5,2.5,0.6,2.9,1.9C81.2,45.7,81.1,47.3,79.7,47.7z M78.2,43.3c-0.8,0.3-0.8,1.2-0.5,2.2 c0.3,1,0.9,1.7,1.7,1.4c0.8-0.3,0.8-1.2,0.5-2.2C79.6,43.8,79,43.1,78.2,43.3z M86,46.3c-1.5,0.2-2.3-1.2-2.4-2.5 c-0.2-1.3,0.3-2.8,1.8-3c1.5-0.2,2.3,1.1,2.5,2.5C88,44.6,87.5,46.1,86,46.3z M85.5,41.7c-0.9,0.1-1.1,1-0.9,2 c0.1,1,0.5,1.9,1.4,1.8c0.9-0.1,1.1-1,0.9-2S86.3,41.6,85.5,41.7z M92.5,46.2c-1.5-0.1-2-1.6-1.9-3c0.1-1.3,0.9-2.7,2.4-2.5 c1.5,0.1,2,1.6,1.9,2.9C94.7,45,93.9,46.4,92.5,46.2z M92.9,41.6c-0.9-0.1-1.2,0.8-1.3,1.8c-0.1,1,0.1,1.9,1,2 c0.9,0.1,1.2-0.8,1.3-1.8C94,42.6,93.8,41.7,92.9,41.6z M98.4,47.4c-0.2-0.1-0.4-0.3-0.3-0.6l1-3.5l-0.8,0.5c-0.1,0.1-0.2,0.1-0.3,0 c-0.2-0.1-0.3-0.3-0.3-0.5c0-0.1,0.1-0.2,0.2-0.2l1.5-0.8c0.1-0.1,0.3-0.1,0.4,0c0.3,0.1,0.5,0.4,0.4,0.7L99,47.1 C98.9,47.3,98.6,47.5,98.4,47.4z M103.1,49.4c-1.3-0.7-1.3-2.2-0.6-3.4c0.6-1.2,1.8-2.2,3.2-1.5c1.3,0.7,1.3,2.2,0.7,3.4 C105.6,49.1,104.4,50,103.1,49.4z M105.2,45.2c-0.8-0.4-1.4,0.3-1.9,1.2c-0.5,0.9-0.6,1.8,0.2,2.2c0.8,0.4,1.4-0.3,1.9-1.2 C105.9,46.5,106,45.6,105.2,45.2z M108.2,52.6c-0.2-0.2-0.2-0.5-0.1-0.7l2.2-2.9l-0.9,0.1c-0.1,0-0.2,0-0.3-0.1 c-0.2-0.1-0.2-0.4-0.1-0.6c0.1-0.1,0.2-0.1,0.3-0.2l1.7-0.2c0.1,0,0.3,0,0.4,0.1c0.2,0.2,0.3,0.5,0.1,0.8l-2.7,3.5 C108.7,52.7,108.4,52.8,108.2,52.6z M111.6,55.8c-0.2-0.2-0.2-0.5,0-0.7l2.7-2.5l-0.9,0c-0.1,0-0.2,0-0.3-0.1 c-0.1-0.2-0.1-0.4,0-0.6c0.1-0.1,0.2-0.1,0.3-0.1l1.7,0c0.1,0,0.3,0.1,0.4,0.2c0.2,0.2,0.2,0.6,0,0.8l-3.2,3 C112.1,56,111.8,56,111.6,55.8z M114.5,59.5c-0.1-0.2-0.1-0.5,0.1-0.7l3-2.1l-0.9-0.2c-0.1,0-0.2-0.1-0.2-0.2 c-0.1-0.2-0.1-0.4,0.1-0.6c0.1-0.1,0.2-0.1,0.3-0.1l1.7,0.3c0.1,0,0.3,0.1,0.4,0.2c0.2,0.2,0.1,0.6-0.1,0.8l-3.6,2.5 C115,59.8,114.7,59.7,114.5,59.5z M116.9,64c-0.6-1.4,0.4-2.5,1.7-3.1c1.2-0.5,2.8-0.5,3.4,0.8c0.6,1.4-0.4,2.5-1.7,3.1 C119.1,65.4,117.5,65.4,116.9,64z M121.2,62.1c-0.3-0.8-1.3-0.7-2.2-0.3c-0.9,0.4-1.6,1.1-1.3,1.8c0.3,0.8,1.3,0.7,2.2,0.3 C120.8,63.6,121.5,62.9,121.2,62.1z M118.8,70.2c-0.3-1.5,1-2.4,2.3-2.6c1.3-0.3,2.8,0.1,3.1,1.6c0.3,1.5-1,2.4-2.3,2.6 C120.6,72,119.1,71.6,118.8,70.2z M123.4,69.2c-0.2-0.8-1.1-1-2.1-0.8c-1,0.2-1.8,0.7-1.6,1.5c0.2,0.8,1.1,1,2.1,0.8 S123.6,70.1,123.4,69.2z"
        />
        <path
          class="data-circle outer"
          d="M134.4,78.1c0-0.3,0.2-0.5,0.5-0.5l3.8,0.1l-0.7-0.7c-0.1-0.1-0.1-0.2-0.1-0.3c0-0.2,0.2-0.4,0.4-0.4c0.1,0,0.2,0,0.3,0.1l1.2,1.3c0.1,0.1,0.2,0.3,0.2,0.4c0,0.3-0.3,0.6-0.6,0.6l-4.6-0.2C134.6,78.6,134.4,78.3,134.4,78.1z M133.8,83.2c0.2-1.5,1.8-2,3.2-1.7c1.4,0.2,2.7,1.1,2.5,2.6c-0.2,1.5-1.8,2-3.2,1.7C134.9,85.6,133.6,84.7,133.8,83.2z M138.6,83.9c0.1-0.9-0.7-1.3-1.8-1.5c-1-0.2-2,0-2.1,0.9c-0.1,0.9,0.7,1.3,1.8,1.5C137.5,85,138.5,84.8,138.6,83.9z M132.5,89.2c0.1-0.3,0.4-0.4,0.6-0.3l3.7,1.1l-0.5-0.8c-0.1-0.1-0.1-0.2,0-0.3c0.1-0.2,0.3-0.3,0.5-0.3c0.1,0,0.2,0.1,0.3,0.2l0.9,1.5c0.1,0.1,0.1,0.3,0,0.4c-0.1,0.3-0.4,0.5-0.7,0.4l-4.4-1.3C132.5,89.8,132.4,89.5,132.5,89.2z M130.9,93.6c0.1-0.2,0.4-0.4,0.6-0.3l3.5,1.5l-0.4-0.9c0-0.1,0-0.2,0-0.3c0.1-0.2,0.3-0.3,0.5-0.2c0.1,0,0.2,0.1,0.2,0.2l0.7,1.6c0.1,0.1,0.1,0.3,0,0.4c-0.1,0.3-0.4,0.4-0.7,0.3l-4.2-1.8C130.9,94.2,130.8,93.9,130.9,93.6z M128.8,97.9c0.1-0.2,0.4-0.3,0.7-0.2l3.3,1.8l-0.3-0.9c0-0.1,0-0.2,0-0.3c0.1-0.2,0.4-0.3,0.6-0.2c0.1,0.1,0.2,0.1,0.2,0.3l0.5,1.7c0,0.1,0,0.3-0.1,0.4c-0.1,0.3-0.5,0.4-0.8,0.2l-4-2.2C128.7,98.4,128.6,98.1,128.8,97.9z M126,102.2c0.9-1.3,2.5-0.9,3.6-0.1c1.1,0.8,1.9,2.2,1,3.5c-0.9,1.3-2.5,1-3.6,0.1C125.9,104.9,125.1,103.5,126,102.2z M129.9,105c0.5-0.7,0-1.5-0.9-2.1s-1.8-0.9-2.3-0.2c-0.5,0.7,0,1.5,0.9,2.1S129.4,105.8,129.9,105z M122.1,107c0.2-0.2,0.5-0.2,0.7,0l2.8,2.6l-0.1-1c0-0.1,0-0.2,0.1-0.3c0.1-0.2,0.4-0.2,0.6,0c0.1,0.1,0.1,0.2,0.1,0.3l0.1,1.8c0,0.1,0,0.3-0.2,0.4c-0.2,0.2-0.6,0.3-0.8,0l-3.4-3.1C121.9,107.5,121.9,107.2,122.1,107z M118.3,110.5c1.2-1,2.6-0.3,3.5,0.8c0.9,1.1,1.3,2.6,0.1,3.6c-1.2,1-2.6,0.3-3.5-0.8C117.5,113,117.1,111.5,118.3,110.5z M121.4,114.2c0.7-0.6,0.3-1.5-0.3-2.3c-0.7-0.8-1.5-1.3-2.2-0.7c-0.7,0.6-0.3,1.5,0.3,2.3C119.9,114.3,120.7,114.8,121.4,114.2z M113.3,114.1c0.2-0.1,0.5-0.1,0.7,0.1l2.1,3.2l0.2-0.9c0-0.1,0.1-0.2,0.2-0.2c0.2-0.1,0.4-0.1,0.6,0.1c0.1,0.1,0.1,0.2,0.1,0.3l-0.4,1.7c0,0.1-0.1,0.3-0.3,0.4c-0.3,0.2-0.6,0.1-0.8-0.2l-2.5-3.9C113,114.6,113.1,114.3,113.3,114.1z M109.2,116.4c0.2-0.1,0.5,0,0.7,0.2l1.7,3.4l0.3-0.9c0-0.1,0.1-0.2,0.2-0.2c0.2-0.1,0.5,0,0.6,0.2c0,0.1,0.1,0.2,0,0.3l-0.5,1.7c0,0.1-0.1,0.3-0.3,0.3c-0.3,0.1-0.6,0-0.8-0.2l-2.1-4.1C108.8,116.9,108.9,116.6,109.2,116.4z M104.4,118.4c1.5-0.5,2.6,0.7,3,2c0.5,1.3,0.3,2.9-1.1,3.4s-2.6-0.6-3-2C102.8,120.5,102.9,118.9,104.4,118.4z M106,122.9c0.8-0.3,0.8-1.3,0.5-2.3s-1-1.8-1.8-1.5c-0.8,0.3-0.8,1.3-0.5,2.3C104.5,122.5,105.1,123.2,106,122.9z M98.5,120.1c0.3-0.1,0.5,0.1,0.6,0.4l0.8,3.7l0.5-0.8c0.1-0.1,0.1-0.1,0.2-0.2c0.2,0,0.4,0.1,0.5,0.3c0,0.1,0,0.2-0.1,0.3l-0.9,1.5c-0.1,0.1-0.2,0.2-0.4,0.2c-0.3,0.1-0.6-0.1-0.7-0.4l-1-4.5C98,120.4,98.2,120.1,98.5,120.1z M93.8,120.8c0.3,0,0.5,0.2,0.5,0.4l0.4,3.8l0.6-0.8c0.1-0.1,0.2-0.1,0.3-0.1c0.2,0,0.4,0.1,0.5,0.4c0,0.1,0,0.2-0.1,0.3l-1.1,1.4c-0.1,0.1-0.2,0.2-0.4,0.2c-0.3,0-0.6-0.2-0.6-0.5l-0.5-4.6C93.3,121,93.5,120.8,93.8,120.8z M89.1,121c0.3,0,0.5,0.2,0.5,0.5l0.1,3.8l0.7-0.7c0.1-0.1,0.2-0.1,0.3-0.1c0.2,0,0.4,0.2,0.4,0.4c0,0.1,0,0.2-0.1,0.3l-1.2,1.3c-0.1,0.1-0.2,0.2-0.4,0.2c-0.3,0-0.6-0.2-0.6-0.5l-0.1-4.6C88.6,121.2,88.8,121,89.1,121z M84.3,120.7c0.3,0,0.5,0.3,0.4,0.5l-0.4,3.8l0.7-0.6c0.1-0.1,0.2-0.1,0.3-0.1c0.2,0,0.4,0.2,0.4,0.5c0,0.1-0.1,0.2-0.2,0.3l-1.4,1.1c-0.1,0.1-0.3,0.1-0.4,0.1c-0.3,0-0.5-0.3-0.5-0.6l0.5-4.6C83.8,120.8,84.1,120.6,84.3,120.7z M79.3,119.7c1.5,0.4,1.8,1.9,1.5,3.3c-0.3,1.4-1.3,2.6-2.8,2.3c-1.5-0.4-1.8-1.9-1.5-3.3C76.8,120.7,77.7,119.4,79.3,119.7z M78.2,124.4c0.9,0.2,1.4-0.6,1.6-1.6c0.2-1,0.1-2-0.7-2.2c-0.9-0.2-1.4,0.6-1.6,1.7S77.3,124.2,78.2,124.4z M73.4,118c0.2,0.1,0.4,0.4,0.3,0.6l-1.3,3.6l0.9-0.4c0.1-0.1,0.2,0,0.3,0c0.2,0.1,0.3,0.3,0.2,0.5c0,0.1-0.1,0.2-0.2,0.2l-1.6,0.7c-0.1,0.1-0.3,0.1-0.4,0c-0.3-0.1-0.5-0.4-0.3-0.7l1.6-4.3C72.8,118,73.1,117.9,73.4,118z M68.7,115.8c1.4,0.7,1.3,2.3,0.6,3.6c-0.6,1.2-1.9,2.2-3.3,1.5c-1.4-0.7-1.3-2.3-0.6-3.6C66,116.1,67.3,115.1,68.7,115.8z M66.5,120.1c0.8,0.4,1.5-0.3,2-1.2c0.5-0.9,0.6-1.9-0.2-2.3c-0.8-0.4-1.5,0.3-2,1.2C65.8,118.7,65.7,119.7,66.5,120.1z M63.4,112.6c0.2,0.2,0.3,0.5,0.1,0.7l-2.2,3.1l0.9-0.2c0.1,0,0.2,0,0.3,0.1c0.2,0.1,0.2,0.4,0.1,0.6c-0.1,0.1-0.2,0.2-0.3,0.2l-1.7,0.3c-0.1,0-0.3,0-0.4-0.1c-0.3-0.2-0.3-0.5-0.2-0.8l2.6-3.8C62.9,112.5,63.2,112.4,63.4,112.6z M59.4,109.3c1.2,1,0.7,2.6-0.3,3.6c-0.9,1-2.4,1.7-3.6,0.6c-1.2-1-0.7-2.6,0.3-3.6C56.8,108.9,58.3,108.3,59.4,109.3z M56.2,112.9c0.7,0.6,1.5,0.1,2.2-0.7c0.7-0.8,1.1-1.7,0.4-2.3c-0.7-0.6-1.5-0.1-2.2,0.7S55.5,112.3,56.2,112.9z M54.9,104.6c1,1.2,0.3,2.6-0.8,3.5c-1.1,0.9-2.6,1.3-3.6,0.1s-0.3-2.6,0.8-3.5C52.3,103.8,53.9,103.4,54.9,104.6z M51.1,107.6c0.6,0.7,1.5,0.3,2.3-0.3c0.8-0.7,1.3-1.5,0.8-2.2c-0.6-0.7-1.5-0.3-2.3,0.3C51.1,106.1,50.6,106.9,51.1,107.6z M51.1,99.2c0.8,1.3-0.1,2.7-1.3,3.4c-1.2,0.7-2.8,0.9-3.6-0.4c-0.8-1.3,0.1-2.7,1.3-3.4C48.6,98,50.2,97.9,51.1,99.2z M46.9,101.7c0.5,0.8,1.4,0.6,2.3,0c0.9-0.5,1.5-1.3,1.1-2.1c-0.5-0.8-1.4-0.6-2.3,0C47.1,100.2,46.4,100.9,46.9,101.7z M48.2,93.7c0.1,0.2,0,0.5-0.2,0.6l-3.5,1.5l0.9,0.3c0.1,0,0.2,0.1,0.2,0.2c0.1,0.2,0,0.5-0.2,0.5c-0.1,0-0.2,0-0.3,0l-1.6-0.6c-0.1,0-0.3-0.2-0.3-0.3c-0.1-0.3,0-0.6,0.3-0.7l4.2-1.9C47.8,93.4,48.1,93.5,48.2,93.7z M46.5,88.9c0.4,1.5-0.8,2.5-2.1,2.9C43,92.2,41.4,92,41,90.5s0.8-2.5,2.1-2.9C44.4,87.2,46,87.4,46.5,88.9z M41.8,90.2c0.3,0.9,1.2,0.9,2.2,0.6c1-0.3,1.8-0.9,1.6-1.7c-0.3-0.9-1.2-0.9-2.2-0.6C42.4,88.8,41.6,89.4,41.8,90.2z M45.1,82.9c0,0.3-0.1,0.5-0.4,0.6l-3.8,0.6l0.8,0.5c0.1,0.1,0.1,0.2,0.2,0.3c0,0.2-0.1,0.4-0.3,0.5c-0.1,0-0.2,0-0.3-0.1l-1.4-1c-0.1-0.1-0.2-0.2-0.2-0.4c0-0.3,0.1-0.6,0.5-0.7l4.5-0.7C44.8,82.4,45,82.6,45.1,82.9z M44.6,77.7c0.1,1.6-1.4,2.3-2.8,2.3c-1.4,0.1-2.9-0.6-2.9-2.1c-0.1-1.6,1.4-2.3,2.8-2.3C43.1,75.6,44.6,76.2,44.6,77.7z M39.8,77.9c0,0.9,1,1.2,2,1.1s2-0.4,1.9-1.3c0-0.9-1-1.2-2-1.1S39.8,77,39.8,77.9z M44.8,71.5c0,0.3-0.3,0.5-0.5,0.4l-3.8-0.3l0.6,0.7c0.1,0.1,0.1,0.2,0.1,0.3c0,0.2-0.2,0.4-0.4,0.4c-0.1,0-0.2-0.1-0.3-0.1l-1.1-1.3c-0.1-0.1-0.1-0.3-0.1-0.4c0-0.3,0.3-0.5,0.6-0.5l4.6,0.4C44.6,71,44.8,71.3,44.8,71.5z M45.6,66.5c-0.3,1.5-1.9,1.9-3.3,1.6c-1.4-0.3-2.7-1.3-2.3-2.8c0.3-1.5,1.9-1.9,3.3-1.6C44.6,64,45.9,65,45.6,66.5z M40.9,65.5c-0.2,0.9,0.6,1.4,1.7,1.6s2,0.1,2.2-0.8c0.2-0.9-0.7-1.4-1.7-1.6C42,64.5,41.1,64.6,40.9,65.5z M47.5,60.1c-0.5,1.5-2.2,1.6-3.5,1.1c-1.3-0.5-2.4-1.6-1.9-3.1c0.5-1.5,2.1-1.6,3.5-1.1C46.9,57.5,48,58.7,47.5,60.1z M43,58.4c-0.3,0.8,0.4,1.5,1.4,1.8s2,0.4,2.3-0.4c0.3-0.8-0.4-1.5-1.4-1.8S43.3,57.6,43,58.4z M50,54.5c-0.1,0.2-0.4,0.3-0.7,0.2L46,52.9l0.3,0.9c0,0.1,0,0.2,0,0.3c-0.1,0.2-0.4,0.3-0.6,0.2c-0.1-0.1-0.2-0.1-0.2-0.3l-0.5-1.7c0-0.1,0-0.3,0-0.4c0.1-0.3,0.5-0.4,0.8-0.2l4.1,2.1C50,54,50.1,54.3,50,54.5z M52.8,50.2c-0.9,1.3-2.5,1-3.6,0.2c-1.1-0.8-2-2.2-1.1-3.5s2.5-1,3.6-0.2C52.8,47.5,53.7,48.9,52.8,50.2z M48.8,47.4c-0.5,0.7,0.1,1.5,0.9,2.1s1.8,0.9,2.3,0.1c0.5-0.7-0.1-1.5-0.9-2.1C50.3,46.9,49.3,46.7,48.8,47.4z M56.6,45.4c-0.2,0.2-0.5,0.2-0.7,0l-2.8-2.6l0.1,1c0,0.1,0,0.2-0.1,0.3c-0.1,0.2-0.4,0.2-0.6,0c-0.1-0.1-0.1-0.2-0.1-0.3L52.2,42c0-0.1,0-0.3,0.2-0.4c0.2-0.2,0.6-0.3,0.8,0l3.4,3.1C56.8,44.8,56.8,45.1,56.6,45.4z M60,42.1c-0.2,0.2-0.5,0.2-0.7,0l-2.5-2.9l0,1c0,0.1-0.1,0.2-0.1,0.3c-0.2,0.1-0.4,0.1-0.6,0C56,40.3,55.9,40.2,56,40l0.1-1.8c0-0.1,0.1-0.3,0.2-0.4c0.2-0.2,0.6-0.2,0.8,0l3,3.4C60.3,41.6,60.2,41.9,60,42.1z M63.7,39.1c-0.2,0.2-0.5,0.1-0.7-0.1l-2.2-3.1l-0.1,0.9c0,0.1-0.1,0.2-0.2,0.3c-0.2,0.1-0.4,0.1-0.6-0.1c-0.1-0.1-0.1-0.2-0.1-0.3l0.3-1.7c0-0.1,0.1-0.3,0.2-0.4c0.2-0.2,0.6-0.1,0.8,0.1l2.7,3.7C64,38.7,64,39,63.7,39.1z M67.8,36.6c-0.2,0.1-0.5,0-0.7-0.2l-1.9-3.3L65,34.1c0,0.1-0.1,0.2-0.2,0.2c-0.2,0.1-0.5,0-0.6-0.2c-0.1-0.1-0.1-0.2,0-0.3l0.5-1.7c0-0.1,0.1-0.3,0.3-0.3c0.3-0.2,0.6-0.1,0.8,0.2l2.3,4C68.1,36.2,68,36.5,67.8,36.6z M72,34.6c-0.2,0.1-0.5,0-0.6-0.3l-1.5-3.5l-0.3,0.9c0,0.1-0.1,0.2-0.2,0.2c-0.2,0.1-0.5,0-0.5-0.2c0-0.1,0-0.2,0-0.3l0.6-1.6c0-0.1,0.2-0.3,0.3-0.3c0.3-0.1,0.6,0,0.7,0.3l1.8,4.2C72.4,34.2,72.3,34.5,72,34.6z M76.9,32.9c-1.5,0.4-2.5-0.8-2.9-2.1c-0.4-1.3-0.2-2.9,1.3-3.4c1.5-0.4,2.5,0.8,2.9,2.1C78.6,30.9,78.4,32.5,76.9,32.9z M75.5,28.3c-0.9,0.2-0.9,1.2-0.6,2.2c0.3,1,0.9,1.8,1.7,1.6c0.9-0.3,0.9-1.2,0.6-2.2C76.9,28.8,76.4,28,75.5,28.3z M82.9,31.5c-0.3,0-0.5-0.1-0.6-0.4l-0.6-3.8l-0.5,0.8c-0.1,0.1-0.2,0.1-0.3,0.2c-0.2,0-0.4-0.1-0.5-0.3c0-0.1,0-0.2,0.1-0.3l1-1.4c0.1-0.1,0.2-0.2,0.4-0.2c0.3,0,0.6,0.1,0.7,0.5l0.7,4.5C83.4,31.2,83.2,31.5,82.9,31.5z M88,31.1c-1.6,0-2.3-1.4-2.3-2.8c0-1.4,0.6-2.9,2.1-2.9c1.6,0,2.3,1.4,2.3,2.8C90.2,29.6,89.6,31.1,88,31.1z M87.9,26.3c-0.9,0-1.2,1-1.2,2c0,1.1,0.4,2,1.3,1.9c0.9,0,1.2-1,1.1-2C89.1,27.2,88.8,26.3,87.9,26.3z M94.2,31.3c-0.3,0-0.5-0.3-0.4-0.5l0.3-3.8l-0.7,0.6c-0.1,0.1-0.2,0.1-0.3,0.1c-0.2,0-0.4-0.2-0.4-0.5c0-0.1,0.1-0.2,0.2-0.3l1.3-1.1c0.1-0.1,0.3-0.1,0.4-0.1c0.3,0,0.5,0.3,0.5,0.6l-0.4,4.6C94.7,31.1,94.5,31.3,94.2,31.3z M99.3,32.2c-1.5-0.3-1.8-1.9-1.5-3.3c0.3-1.4,1.3-2.6,2.8-2.3c1.5,0.3,1.9,1.9,1.6,3.3C101.8,31.2,100.8,32.5,99.3,32.2z M100.3,27.4c-0.9-0.2-1.4,0.6-1.6,1.7c-0.2,1-0.1,2,0.8,2.2c0.9,0.2,1.4-0.6,1.6-1.7C101.3,28.6,101.2,27.6,100.3,27.4z M105.2,33.9c-0.2-0.1-0.4-0.4-0.3-0.6l1.3-3.6l-0.9,0.4c-0.1,0.1-0.2,0.1-0.3,0c-0.2-0.1-0.3-0.3-0.2-0.5c0-0.1,0.1-0.2,0.2-0.2l1.6-0.8c0.1-0.1,0.3-0.1,0.4,0c0.3,0.1,0.5,0.4,0.4,0.7l-1.5,4.3C105.8,33.8,105.5,34,105.2,33.9z M109.9,36c-1.4-0.7-1.3-2.3-0.7-3.6c0.6-1.2,1.9-2.2,3.3-1.5c1.4,0.7,1.3,2.3,0.7,3.6C112.6,35.7,111.3,36.7,109.9,36z M112.1,31.7c-0.8-0.4-1.5,0.3-2,1.2c-0.5,0.9-0.6,1.9,0.2,2.3c0.8,0.4,1.5-0.3,2-1.2S112.9,32.1,112.1,31.7z M115.2,39.1c-0.2-0.1-0.3-0.5-0.1-0.7l2.2-3.1l-0.9,0.2c-0.1,0-0.2,0-0.3-0.1c-0.2-0.1-0.2-0.4-0.1-0.6c0.1-0.1,0.2-0.2,0.3-0.2l1.7-0.3c0.1,0,0.3,0,0.4,0.1c0.3,0.2,0.3,0.5,0.2,0.8l-2.6,3.8C115.8,39.2,115.4,39.3,115.2,39.1z M119.2,42.4c-1.2-1-0.7-2.6,0.2-3.6c0.9-1,2.4-1.7,3.5-0.7c1.2,1,0.7,2.6-0.2,3.6C121.9,42.7,120.4,43.4,119.2,42.4z M122.4,38.7c-0.7-0.6-1.5-0.1-2.2,0.7c-0.7,0.8-1.1,1.7-0.4,2.3c0.7,0.6,1.5,0.1,2.2-0.7C122.7,40.2,123.1,39.3,122.4,38.7z M123.6,46.7c-0.2-0.2-0.2-0.5,0-0.7l2.9-2.5l-1,0c-0.1,0-0.2-0.1-0.3-0.1c-0.1-0.2-0.1-0.4,0-0.6c0.1-0.1,0.2-0.1,0.3-0.1l1.8,0.1c0.1,0,0.3,0.1,0.4,0.2c0.2,0.2,0.2,0.6,0,0.8l-3.4,3C124.1,46.9,123.8,46.9,123.6,46.7z M126.7,50.8c-0.9-1.3,0-2.7,1.1-3.4c1.2-0.8,2.7-1.1,3.6,0.2c0.9,1.3,0,2.7-1.1,3.4C129.2,51.9,127.6,52.1,126.7,50.8z M130.7,48.1c-0.5-0.7-1.4-0.5-2.3,0.1c-0.9,0.6-1.5,1.4-1,2.1c0.5,0.7,1.4,0.5,2.3-0.1C130.6,49.6,131.2,48.9,130.7,48.1z M129.8,56.1c-0.1-0.2,0-0.5,0.2-0.7l3.4-1.7l-0.9-0.3c-0.1,0-0.2-0.1-0.2-0.2c-0.1-0.2,0-0.5,0.2-0.6c0.1,0,0.2-0.1,0.3,0l1.7,0.5c0.1,0,0.3,0.1,0.3,0.3c0.1,0.3,0,0.6-0.2,0.8l-4.1,2.1C130.3,56.5,130,56.4,129.8,56.1z M131.7,60.5c-0.1-0.2,0-0.5,0.3-0.6l3.6-1.4l-0.9-0.4c-0.1,0-0.2-0.1-0.2-0.2c-0.1-0.2,0-0.5,0.2-0.5c0.1,0,0.2,0,0.3,0l1.6,0.7c0.1,0.1,0.2,0.2,0.3,0.3c0.1,0.3,0,0.6-0.3,0.7l-4.3,1.6C132.1,60.9,131.8,60.8,131.7,60.5z M133.2,65.4c-0.4-1.5,0.9-2.5,2.3-2.8c1.4-0.3,2.9,0,3.3,1.5c0.4,1.5-0.9,2.5-2.3,2.8C135.1,67.3,133.5,67,133.2,65.4z M137.9,64.3c-0.2-0.9-1.2-1-2.2-0.7s-1.9,0.8-1.6,1.6c0.2,0.9,1.2,1,2.2,0.7C137.2,65.7,138.1,65.2,137.9,64.3z M134.2,71.9c-0.1-1.5,1.3-2.3,2.7-2.5c1.4-0.1,2.9,0.4,3,2s-1.2,2.3-2.6,2.5C135.9,74,134.4,73.5,134.2,71.9z M139,71.5c-0.1-0.9-1-1.1-2.1-1c-1,0.1-1.9,0.5-1.9,1.4c0.1,0.9,1,1.1,2.1,1C138.2,72.8,139.1,72.4,139,71.5z"
        />
        <circle class="glow-circle" cx="97.5" cy="76" r="21.5" />
        <g class="glow-lines">
          <line x1="97.5" y1="43" x2="97.5" y2="7" />
          <line x1="74.5" y1="52.5" x2="49.1" y2="27.1" />
          <line x1="74.5" y1="98.5" x2="49.1" y2="123.9" />
          <line x1="97.5" y1="108" x2="97.5" y2="144" />
          <line x1="120.5" y1="98.5" x2="145.9" y2="123.9" />
          <line x1="120.5" y1="52.5" x2="145.9" y2="27.1" />
        </g>
        <g class="key-body">
          <path
            class="key-border"
            d="M172,74v4.5c0,0.3-0.2,0.5-0.5,0.5H168v17.5c0,0.3-0.2,0.5-0.5,0.5h-6c-0.3,0-0.5-0.2-0.5-0.5V79h-5v14.5c0,0.3-0.2,0.5-0.5,0.5h-6 c-0.3,0-0.5-0.2-0.5-0.5V79h-41v1c0,1.2-1.1,2.2-2.3,2c-1-0.2-1.7-1.1-1.7-2.1V79H67v-7h37v-1c0-1.2,1.1-2.2,2.3-2 c1,0.2,1.7,1.1,1.7,2.1V72h62C171.1,72,172,72.9,172,74z M57,76c0,13.8-11.2,25-25,25S7,89.8,7,76s11.2-25,25-25S57,62.2,57,76z M49,76c0-9.4-7.6-17-17-17s-17,7.6-17,17s7.6,17,17,17S49,85.4,49,76z M62,85.5v-20c0-1.4-1.1-2.5-2.5-2.5h0 c-1.4,0-2.5,1.1-2.5,2.5v20c0,1.4,1.1,2.5,2.5,2.5h0C60.9,88,62,86.9,62,85.5z M67,80.5v-10c0-1.4-1.1-2.5-2.5-2.5l0,0 c-1.4,0-2.5,1.1-2.5,2.5v10c0,1.4,1.1,2.5,2.5,2.5l0,0C65.9,83,67,81.9,67,80.5z"
          />
          <path
            class="key-highlight"
            d="M32,97c-11.6,0-21-9.4-21-21c0-11.6,9.4-21,21-21s21,9.4,21,21C53,87.6,43.6,97,32,97z M32,58 c-9.9,0-18,8.1-18,18c0,9.9,8.1,18,18,18s18-8.1,18-18C50,66.1,41.9,58,32,58z M107,73v-2c0-0.6-0.4-1-1-1s-1,0.4-1,1v2H70v2h99v-2 H107z M59.5,64c-0.8,0-1.5,0.7-1.5,1.5v2c0-0.8,0.7-1.5,1.5-1.5s1.5,0.7,1.5,1.5v-2C61,64.7,60.3,64,59.5,64z M64.5,69 c-0.8,0-1.5,0.7-1.5,1.5v2c0-0.8,0.7-1.5,1.5-1.5s1.5,0.7,1.5,1.5v-2C66,69.7,65.3,69,64.5,69z"
          />
          <path
            class="key-darkness"
            d="M167,78v16v2h-2h-3v-2h3V78H167z M153,78v13h-3v2h3h2v-2V78H153z M169,73v5h2v-4c0-0.6-0.4-1-1-1H169z M68,73v5h2v-5H68z M59,86.9V64.1c-0.6,0.2-1,0.8-1,1.4v20C58,86.2,58.4,86.7,59,86.9z M64,81.9V69.1c-0.6,0.2-1,0.8-1,1.4v10C63,81.2,63.4,81.7,64,81.9z"
          />
        </g>
        <g class="stars">
          <path
            d="M91.9,44c-3.8,2.2-8.7,0.9-10.9-2.9l-1-1.7l1,1.7c2.2,3.8,0.9,8.7-2.9,10.9l-1.7,1l1.7-1c3.8-2.2,8.7-0.9,10.9,2.9l1,1.7l-1-1.7 C86.8,51.1,88.1,46.2,91.9,44l1.7-1L91.9,44z"
          />
          <path
            d="M134.9,92c-3.8,2.2-8.7,0.9-10.9-2.9l-1-1.7l1,1.7c2.2,3.8,0.9,8.7-2.9,10.9l-1.7,1l1.7-1c3.8-2.2,8.7-0.9,10.9,2.9l1,1.7l-1-1.7 C129.8,99.1,131.1,94.2,134.9,92l1.7-1L134.9,92z"
          />
          <path
            d="M66.9,114c-3.8,2.2-8.7,0.9-10.9-2.9l-1-1.7l1,1.7c2.2,3.8,0.9,8.7-2.9,10.9l-1.7,1l1.7-1c3.8-2.2,8.7-0.9,10.9,2.9l1,1.7l-1-1.7 C61.8,121.1,63.1,116.2,66.9,114l1.7-1L66.9,114z"
          />
        </g>
      </svg>
    </div>
    <div class="splash-message">{{ message }}</div>
  </div>
</template>

<script setup lang="ts">
export interface Props {
  glow: boolean;
  message: string;
}

defineProps<Props>();
</script>

<style scoped>
.splash-page {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.key {
  flex: none;
  width: 90px;
  height: 95px;
}

.splash-message {
  flex: none;
  font-style: italic;
}

/* Start: Key */
@keyframes jump {
  0%,
  20%,
  53%,
  80%,
  100% {
    transition-timing-function: cubic-bezier(0.215, 0.61, 0.355, 1);
    transform: translate3d(0, 0, 0);
  }

  40%,
  43% {
    transition-timing-function: cubic-bezier(0.755, 0.05, 0.855, 0.06);
    transform: translate3d(0, -16px, 0) rotate(8deg);
  }

  70% {
    transition-timing-function: cubic-bezier(0.755, 0.05, 0.855, 0.06);
    transform: translate3d(0, -6px, 0) rotate(-2deg);
  }

  90% {
    transform: translate3d(0, -3px, 0) rotate(2deg);
  }
}

@keyframes clockRotate {
  0% {
    transform: rotate(0deg);
  }

  100% {
    transform: rotate(360deg);
  }
}

@keyframes glowStarD0 {
  0% {
    transform: rotate(0deg) scale3d(0.5, 0.5, 1);
    opacity: 0;
  }

  10% {
    opacity: 1;
  }

  20% {
    opacity: 1;
  }

  40%,
  100% {
    transform: rotate(90deg) scale3d(1.2, 1.2, 1);
    opacity: 0;
  }
}

@keyframes glowStarD1 {
  0%,
  15% {
    transform: rotate(0deg) scale3d(0.5, 0.5, 1);
    opacity: 0;
  }

  25% {
    opacity: 1;
  }

  35% {
    opacity: 1;
  }

  55%,
  100% {
    transform: rotate(90deg) scale3d(1.2, 1.2, 1);
    opacity: 0;
  }
}

@keyframes glowStarD2 {
  0%,
  30% {
    transform: rotate(0deg) scale3d(0.5, 0.5, 1);
    opacity: 0;
  }

  40% {
    opacity: 1;
  }

  50% {
    opacity: 1;
  }

  70%,
  100% {
    transform: rotate(90deg) scale3d(1.2, 1.2, 1);
    opacity: 0;
  }
}

@keyframes glowCircle {
  0% {
    stroke-dashoffset: -10;
  }

  100% {
    stroke-dashoffset: 10;
  }
}

@keyframes glowLineD0 {
  0% {
    stroke-dashoffset: 40;
  }

  45%,
  100% {
    stroke-dashoffset: -40;
  }
}

@keyframes glowLineD1 {
  0%,
  5% {
    stroke-dashoffset: 40;
  }

  50%,
  100% {
    stroke-dashoffset: -40;
  }
}

@keyframes glowLineD2 {
  0%,
  10% {
    stroke-dashoffset: 40;
  }

  55%,
  100% {
    stroke-dashoffset: -40;
  }
}

@keyframes glowLineD3 {
  0%,
  15% {
    stroke-dashoffset: 40;
  }

  60%,
  100% {
    stroke-dashoffset: -40;
  }
}

@keyframes glowLineD4 {
  0%,
  20% {
    stroke-dashoffset: 40;
  }

  65%,
  100% {
    stroke-dashoffset: -40;
  }
}

@keyframes glowLineD5 {
  0%,
  25% {
    stroke-dashoffset: 40;
  }

  70%,
  100% {
    stroke-dashoffset: -40;
  }
}

.key {
  width: 180px;
  height: 150px;
}

.key svg {
  width: 180px;
  height: 150px;
  stroke-linecap: round;
  stroke-linejoin: round;
}

.key svg .key-body {
  transform-origin: 80px 78px;
}

.key svg .key-body .key-border {
  fill: #fff;
  stroke: #9aa8b1;
  stroke-width: 2;
  transition: all 0.4s;
}

.key svg .key-body .key-highlight {
  fill: #e8ebed;
  transition: all 0.4s;
}

.key svg .key-body .key-darkness {
  fill: #e8ebed;
  transition: all 0.4s;
}

.key svg .data-circle {
  transform-origin: 89.5px 76px;
  transition: fill 0.8s;
  transition-delay: 0.2s;
  animation: clockRotate 18s linear 0s infinite reverse;
  fill: rgba(137, 148, 155, 0.3);
}

.key svg .data-circle.inner {
  animation: clockRotate 20s linear 0s infinite;
}

.key svg .glow-circle {
  opacity: 0;
  fill: none;
  stroke: #cb9916;
  stroke-width: 3;
  stroke-dasharray: 30, 14, 0.1, 14;
  transition: all 0.4s;
  animation: glowCircle 3s ease-in-out 0s infinite alternate;
}

.key svg .glow-lines {
  opacity: 0;
  transition: all 0.4s;
}

.key svg .glow-lines line {
  stroke: #cb9916;
  stroke-width: 3;
  stroke-dasharray: 40 100;
  stroke-dashoffset: 40;
  animation: glowLineD0 2s linear 0s infinite;
  animation-play-state: paused;
}

.key svg .glow-lines line:nth-child(1) {
  animation-name: glowLineD0;
}

.key svg .glow-lines line:nth-child(2) {
  animation-name: glowLineD3;
}

.key svg .glow-lines line:nth-child(3) {
  animation-name: glowLineD1;
}

.key svg .glow-lines line:nth-child(4) {
  animation-name: glowLineD4;
}

.key svg .glow-lines line:nth-child(5) {
  animation-name: glowLineD2;
}

.key svg .glow-lines line:nth-child(6) {
  animation-name: glowLineD5;
}

.key svg .stars {
  opacity: 0;
  transition: opacity 0.4s;
}

.key svg .stars path {
  fill: #fff;
  stroke: #cb9916;
  stroke-width: 2;
  transform-origin: 50% 50%;
  animation: glowStarD0 2s linear 0.9s infinite;
  animation-play-state: paused;
  transform: rotate(0deg) scale3d(0.5, 0.5, 1);
  opacity: 0;
}

.key svg .stars path:nth-child(1) {
  animation-name: glowStarD0;
  transform-origin: 87px 47px;
}

.key svg .stars path:nth-child(2) {
  animation-name: glowStarD1;
  transform-origin: 128px 96px;
}

.key svg .stars path:nth-child(3) {
  animation-name: glowStarD2;
  transform-origin: 62px 117px;
}

.key.on svg .key-body {
  animation: jump 0.7s;
}

.key.on svg .key-body .key-border {
  fill: #ffd055;
  stroke: #cb9916;
  transition-delay: 0.2s;
}

.key.on svg .key-body .key-highlight {
  fill: #ffe155;
  transition-delay: 0.2s;
}

.key.on svg .key-body .key-darkness {
  fill: #f1bf3c;
  transition-delay: 0.2s;
}

.key.on svg .data-circle {
  fill: rgba(137, 148, 155, 0);
  transition-duration: 0.4s;
  transition-delay: 0;
}

.key.on svg .glow-circle {
  opacity: 0.15;
  transition-delay: 0.2s;
}

.key.on svg .glow-lines {
  opacity: 0.15;
  transition-delay: 0.2s;
}

.key.on svg .glow-lines line {
  animation-play-state: running;
}

.key.on svg .stars {
  opacity: 1;
  transition-delay: 0.2s;
}

.key.on svg .stars path {
  animation-play-state: running;
}
</style>
