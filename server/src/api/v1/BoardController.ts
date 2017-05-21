import {Get, Post, JsonController, Param, Body, Req, UseBefore} from 'routing-controllers';
import {JSONWebToken} from '../../utils/JSONWebToken';
import {Thing} from '../../entities/thing.model';
import {UserAuthenticatorMiddleware} from '../../middleware/UserAuthenticatorMiddleware';
import {BlockchainClient} from '../../blockchain/client/blockchainClient';
import {Service} from 'typedi';

@JsonController('/board')
//@UseBefore(UserAuthenticatorMiddleware)
@Service()
export class ThingsController {
  public constructor(private blockchainClient: BlockchainClient) { }

 @Get('/')
  public post(@Body() thing: Thing, @Req() request: any): any {
    let enrollmentID = new JSONWebToken(request).getUserID();
    return this.blockchainClient.query('getCurrentState', [], 'player1');
  }
}