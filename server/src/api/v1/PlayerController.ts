import {Get, Post, JsonController, Param, Body, Req, UseBefore} from 'routing-controllers';
import {JSONWebToken} from '../../utils/JSONWebToken';
import {Thing} from '../../entities/thing.model';
import {UserAuthenticatorMiddleware} from '../../middleware/UserAuthenticatorMiddleware';
import {BlockchainClient} from '../../blockchain/client/blockchainClient';
import {Service} from 'typedi';

@JsonController('/player')
//@UseBefore(UserAuthenticatorMiddleware)
@Service()
export class PlayerController {
  public constructor(private blockchainClient: BlockchainClient) { }

 @Get('/roll')
  public roll(@Body() body: any, @Req() request: any): any {
    return this.blockchainClient.invoke('rollDice', [], 'player1');
  }
  @Get('/roll2')
  public roll2(@Body() body: any, @Req() request: any): any {
    return this.blockchainClient.invoke('rollDice', [], 'player2');
  }
 @Get('/action')
  public action(@Body() body: any, @Req() request: any): any {
    return this.blockchainClient.invoke('playeraction', ['buy'], 'player1');
  }
  @Get('/action2')
  public action2(@Body() body: any, @Req() request: any): any {
    return this.blockchainClient.invoke('playeraction', ['buy'], 'player2');
  }
 @Get('/start')
  public start(@Body() body: any, @Req() request: any): any {
    return this.blockchainClient.invoke('playerstart', [], 'player1');
  }
}
