package controller

import (
    "net/http"
    "github.com/gorilla/context"
    "github.com/gorilla/mux"
    "kalaxia-game-api/exception"
    "kalaxia-game-api/manager"
    "kalaxia-game-api/model"
    "kalaxia-game-api/utils"
    "strconv"
)

func GetPlanet(w http.ResponseWriter, r *http.Request) {
    /**
     * return the planet (may return partial information) from http request.
     * if the player is the owner of the planet all the information is there
     * but if the player is not the owner the information is partial
     */
    player := context.Get(r, "player").(*model.Player)
    id, _ := strconv.ParseUint(mux.Vars(r)["id"], 10, 16)

    utils.SendJsonResponse(w, 200, manager.GetPlanet(uint16(id), player.Id))
}

func UpdatePlanetSettings(w http.ResponseWriter, r *http.Request) {
    /**
     * request from http to update setting.
     * Parse data form the request.
     * Check is the player is the owner of the planet before updating.
     */
    player := context.Get(r, "player").(*model.Player)

    id, _ := strconv.ParseUint(mux.Vars(r)["id"], 10, 16)
    // mux.Vars get the id associate with the route var id
    planet := manager.GetPlanet(uint16(id), player.Id)

    if player.Id != planet.Player.Id {
        panic(exception.NewHttpException(http.StatusForbidden, "", nil))
    }
    data := utils.DecodeJsonRequest(r)
    settings := &model.PlanetSettings{
        ServicesPoints: uint8(data["services_points"].(float64)),
        BuildingPoints: uint8(data["building_points"].(float64)),
        MilitaryPoints: uint8(data["military_points"].(float64)),
        ResearchPoints: uint8(data["research_points"].(float64)),
    }
    manager.UpdatePlanetSettings(planet, settings)
    utils.SendJsonResponse(w, 200, settings)
}
