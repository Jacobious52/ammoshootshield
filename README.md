# _Ammo Shoot Shield_
_Ammo Shoot Shield_ competitive programming enginer.

## Game
Each round each player makes 1 move:
- ammo: Increase ammo by 1
- shoot: Shoot at other player. Ammo must be > 0. Consumes 1 ammo.
- shield: Block a shot from the other player.

If a player shoots and the other player shields, continue playing.  
If a player shoots with 5 or more ammo, and the other player shields, player shooting wins.  
If a players shoots and the other player ammos, player shooting wins.  
If both players shoot the player with the most ammo wins.  

## Competition Objective
Create an I/O program in any language or implement a Go package implementing the Player interface to create your own AI to compete against other AIs.
