import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Team, TeamService, Test, TestService, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInTeacherController } from '../shared/app.controller';
import { Examen } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-test-teams',
  templateUrl: './test-teams.component.html',
  styleUrls: ['./test-teams.component.css']
})
export class TestTeamsComponent extends LoggedInTeacherController implements OnInit {

  teams : Team[]
  test: Test
  routeSub: Subscription
  id: number
  addTeamTeamname: string
  kickingTeamname: string
  isInAdminTeam: boolean
  autoteams: Team[]
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private tS: TestService, private route: ActivatedRoute,
      private teamS: TeamService) {
    super(session, router, data, userS)
    this.test = new Examen()
    this.teams = []
    this.id = 0
    this.autoteams = []
    this.isInAdminTeam = false
    this.addTeamTeamname = ""
    this.kickingTeamname = ""
    this.routeSub = this.route.params.subscribe(params => {
      this.id = params['testid']
      this.borrarMensaje()
      this.getTest(true)
    });
    this.changeGetAutoTeams()
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    super.onDestroy()
    this.borrarMensaje()
    this.routeSub.unsubscribe()
  }

  doHasUserAction() {
    if (this.id!=undefined && this.id != 0)
      this.getIsInAdminTeam(true)
  }

  getTest(primera: boolean) {
    this.tS.getTest(this.id).subscribe(
      resp => {
        this.test = Examen.constructorFromTest(resp)
        this.getTeamsTest(true)
        if (!this.getSessionUser().isEmpty())
          this.getIsInAdminTeam(true)
      },
      err => this.handleErrRelog(err, "obtener test", primera, this.getTest, this)
    )
  }

  getTeamsTest(primera: boolean){
    this.tS.getAdminTeamsFromTest(this.id).subscribe(
      resp =>{
        this.teams = resp
      },
      err => this.handleErrRelog(err, "obtener equipos del test", primera, this.getTeamsTest, this)
    )
  }

  getIsInAdminTeam(primera: boolean) {
    this.userS.getSharedTestFromUser(this.getSessionUser().getUsername(), this.id).subscribe(
      resp => this.isInAdminTeam = true,
      err => {
        if(err.status!=410)
          this.handleErrRelog(err, "saber si el usuario administra el test", primera, this.getIsInAdminTeam, this)
      }
    )
  }

  isPermisosAdministracion() : boolean{
    return (this.getSessionUser().getUsername() == this.test.username) || this.isInAdminTeam
  }

  checkPermisosAdministracion(): boolean {
    return (this.getSessionUser().isAdmin() || this.isPermisosAdministracion()) // && this.test.editable 
  }

  addTeamSubmit(){
    this.addTeam(true)
  }

  addTeam(primera: boolean){
    this.tS.addAdminTeamToTest(this.addTeamTeamname, this.id).subscribe(
      resp => this.getTeamsTest(true),
      err => this.handleErrRelog(err, "añadir equipo a un test", primera, this.addTeam, this)
    )
  }

  kickTeam(teamname: string){
    this.kickingTeamname = teamname
    this.kickT(true)
  }

  kickT(primera: boolean){
    this.tS.removeAdminTeamToTest(this.kickingTeamname, this.id).subscribe(
      resp => this.getTeamsTest(true),
      err => this.handleErrRelog(err, "eliminar equipo de un test", primera, this.kickT, this)
    )
  }

  changeGetAutoTeams(){
    this.getAutoTeams(true)
  }

  getAutoTeams(primera: boolean){
    this.teamS.getTeams(this.addTeamTeamname, undefined, 20).subscribe(
      resp=>{
        this.autoteams=resp
      },
      err => this.handleErrRelog(err, "obtener equipos que empiezan por ese teamname", primera, this.getAutoTeams, this)
    )
  }

}
