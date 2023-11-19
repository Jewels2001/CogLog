<template>
  <div>
    <h1>
        <slot name = "title">
          <img alt="Cerecode logo" class="logo" src="@/assets/cclogocolor.svg" width="525" height="125" />
        </slot>
    </h1>
  
    <div class="item">
    <!--
    <i><slot name="icon"></slot></i>
    -->
      <div class="details">
        <h3>
          <slot name="heading"></slot>
        </h3>
        <slot></slot>
      </div>
    </div>
  </div>
</template>

<style scoped>
.logo{
  position: absolute;
  top: 25%;
  left: 50%;
  transform: translate(-50%, -50%);
}
/*
h1{
  font-size: 5rem;
  font-weight: 500;
  color: var(--color-heading);
  position: absolute;
  top: -8600%;
  left: 50%;
  transform: translate(-50%, -50%);
}
*/
.item {
  text-align: center;
  display: flex;
  position: absolute;
  top: 75%;
  left: 50%;
  transform: translate(-50%, -50%);
}

.right_item {
  text-align: right;
  display: flex;
  position: relative;
}

.left_item {
  text-align: left;
  display: flex;
  position: relative;
}

.details {
  flex: 1;
}

/*
i {
  display: flex;
  place-items: center;
  place-content: center;
  width: 32px;
  height: 32px;
  color: var(--color-text);
}
*/

h3 {
  font-size: 1.2rem;
  font-weight: 500;
  margin-bottom: 0.4rem;
  color: var(--color-heading);
}

/*
@media (min-width: 1024px) {
  .item {
    margin-top: 0;
    padding: 0.4rem 0 1rem calc(var(--section-gap) / 2);
  }

  i {
    top: calc(50% - 25px);
    left: -26px;
    position: absolute;
    border: 1px solid var(--color-border);
    background: var(--color-background);
    border-radius: 8px;
    width: 50px;
    height: 50px;
  }

  .item:before {
    content: ' ';
    border-left: 1px solid var(--color-border);
    position: absolute;
    left: 0;
    bottom: calc(50% + 25px);
    height: calc(50% - 25px);
  }

  .item:after {
    content: ' ';
    border-left: 1px solid var(--color-border);
    position: absolute;
    left: 0;
    top: calc(50% + 25px);
    height: calc(50% - 25px);
  }

  .item:first-of-type:before {
    display: none;
  }

  .item:last-of-type:after {
    display: none;
  }
}
*/
</style>
