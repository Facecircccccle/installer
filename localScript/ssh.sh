#!/usr/bin/expect
#./ssh.sh root 123456 192.168.1.1
set timeout 10
set username [lindex $argv 0]
set password [lindex $argv 1]
set hostname [lindex $argv 2]
spawn ssh-copy-id -i /root/.ssh/id_rsa.pub $username@$hostname
expect {
            #first connect, no public key in ~/.ssh/known_hosts
            "Are you sure you want to continue connecting (yes/no)?" {
            send "yes\r"
            expect "password:"
                send "$password\r"
                "Permission denied, please try again."{
                send \x03
            }
            }
            #already has public key in ~/.ssh/known_hosts
            "password:" {
                send "$password\r"
                "Permission denied, please try again."{
                send \x03
            }
            }

            "Now try logging into the machine" {
                #it has authorized, do nothing!
            }
        }
expect eof
