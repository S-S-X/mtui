export const get = name => fetch(`api/playerinfo/${name}`).then(r => r.json());