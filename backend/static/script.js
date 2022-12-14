const base = window.location.href;

const addUrl = async (e) => {
  const form = document.getElementById("form");
  e.preventDefault();
  const formData = new FormData(form);
  let data = {};
  for (let key of formData.keys()) {
    data[key] = formData.get(key);
  }
  if (!/^https?:\/\//i.test(data.long_url)) {
    data.long_url = "http://" + data.long_url;
  }
  const status = document.getElementById("status");
  status.innerHTML = "Loading...";
  const response = await fetch(`${base}api/link/create`, {
    method: "POST",
    body: JSON.stringify(data),
  });

  
  if (response.status === 200) {
    status.innerHTML = `url created! you can open using link <a target="_blank" href="${base}${data.short_url}">${base}${data.short_url} </a>`;
  } else if (response.status === 409) {
    let data = await response.json();
    console.log(data);
    status.innerHTML = "Short url already exist. Please use other url";
  } else if (response.status === 400) {
    status.innerHTML = "link format is wrong. please fix them";
  } else {
    status.innerHTML = "something is wrong, try again later";
  }
};
