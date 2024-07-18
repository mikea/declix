# declix

Declarative Linux (declix) is a file format
that describes the state of Linux resources
and a command line tool that applies the
description to a running Linux system.

Declix doesn't try to manage _full_ system
configuration but concerns itself only with
resources declared. Thus you can use Declix
to manage only the part of the system, to
take over existing one or live side-by-side
with other configuration management systems.

Declix can synchronize system state locally
or remotely using ssh. Bash and coreutils
are the only system dependencies required
to be present on a target computer. Declix
doesn't have any persistent state. It sends
all resources needed to bring the target
up-to-date and sends them over ssh connection,
thus allowing deployments in a restricted
environment.

## Example

### Define Target

Define a target in `declpi-target.pkl`:

```pkl
amends "../target/Target.pkl"

target = new SshConfig {
    user = "mike"
    address = "declpi:22"
    privateKey = "/home/mike/.ssh/id_ed25519"
}

```

### Install Packages

Define system configuration in `declpi.pkl`.
First let's make sure that all the packages that
we need are installed:

```pkl
amends "../resources/Resources.pkl"

import "../resources/apt/Apt.pkl"

resources = new Listing {
    ...Apt.install(new Listing {
        "tmux" "vim""git"  "mc" "nnn" "ncdu"
        "btop" "htop" "iftop"
        "kitty-terminfo"
        "ssh" "syncthing"
    })
}
```

We can evaluate the config by running
`pkl eval local/declpi.pkl`
to see what does it expand into:

```pkl
resources {
  new {
    type = "apt"
    name = "tmux"
    state = "installed"
    updateBeforeInstall = false
  }
  new {
    type = "apt"
    name = "vim"
    state = "installed"
    updateBeforeInstall = false
  }
  new {
    type = "apt"
    name = "git"
    state = "installed"
    updateBeforeInstall = false
  }
  new {
    type = "apt"
    name = "mc"
    state = "installed"
    updateBeforeInstall = false
  }
  new {
    type = "apt"
    name = "nnn"
    state = "installed"
    updateBeforeInstall = false
  }
  new {
    type = "apt"
    name = "ncdu"
    state = "installed"
    updateBeforeInstall = false
  }
  new {
    type = "apt"
    name = "btop"
    state = "installed"
    updateBeforeInstall = false
  }
  new {
    type = "apt"
    name = "htop"
    state = "installed"
    updateBeforeInstall = false
  }
  new {
    type = "apt"
    name = "iftop"
    state = "installed"
    updateBeforeInstall = false
  }
  new {
    type = "apt"
    name = "kitty-terminfo"
    state = "installed"
    updateBeforeInstall = false
  }
  new {
    type = "apt"
    name = "ssh"
    state = "installed"
    updateBeforeInstall = false
  }
  new {
    type = "apt"
    name = "syncthing"
    state = "installed"
    updateBeforeInstall = false
  }
}
```

From here we can check the state of the target system:

```
$ go run main.go state -t local/declix-target.pkl -r local/declix.pkl
Target: &{declix:22 mike /home/mike/.ssh/id_ed25519}
Checking...
Resource Id        | Current State     | Expected State                                                               
-------------------------------------------------------
apt:tmux           | missing           | installed
apt:vim            | missing           | installed
apt:git            | missing           | installed
apt:mc             | missing           | installed
apt:nnn            | missing           | installed
apt:ncdu           | missing           | installed
apt:btop           | missing           | installed
apt:htop           | missing           | installed
apt:iftop          | missing           | installed
apt:kitty-terminfo | missing           | installed
apt:ssh            | 1:9.2p1-2+deb12u2 | installed
apt:syncthing      | missing           | installed
```

And we can update the system:

```
$ go run main.go apply -t local/declix-target.pkl -r local/declix.pkl
✓ +apt:tmux                                                                                                           
✓ +apt:vim                                                                                                            
✓ +apt:git                                                                                                            
✓ +apt:mc                                                                                                             
✓ +apt:nnn                                                                                                            
✓ +apt:ncdu                                                                                                           
✓ +apt:btop                                                                                                           
✓ +apt:htop                                                                                                           
✓ +apt:iftop                                                                                                          
✓ +apt:kitty-terminfo                                                                                                 
✓ +apt:syncthing                                                                                                      
```

#### Defining Modules

We can define modules to encapsulate more complicated scenarious.

E.g. installing packages from a different apt source.

`zerotier.pkl`:

```pkl
import "../resources/filesystem/FileSystem.pkl" as FS
import "../resources/apt/Apt.pkl"
import "../resources/Resources.pkl"
import "../content/Content.pkl"

url = "https://download.zerotier.com/"
codename = "bookworm"

function install(): Listing<Resources.Resource> = new Listing {
    new FS.File {
        path = "/etc/apt/trusted.gpg.d/zerotier.com.gpg"
        content = new Content.File {
            file = "modules/zerotier.com.gpg"
        }
        owner = "root"
        group = "root"
        permissions = "644"
    }

    new FS.File {
        path = "/etc/apt/sources.list.d/zerotier.list"
        owner = "root"
        group = "root"
        permissions = "644"
        content = "deb [signed-by=/etc/apt/trusted.gpg.d/zerotier.com.gpg] \(url)debian/\(codename) \(codename) main"
    }

    new Apt.Package {
        name = "zerotier-one"
        state = "installed"
        updateBeforeInstall = true
    }
}

function remove() : Listing<Resources.Resource> = new Listing {
    new Apt.Package {
        name = "zerotier-one"
        state = "missing"
        updateBeforeInstall = true
    }
    new FS.File {
        path = "/etc/apt/sources.list.d/zerotier.list"
        state = new FS.Missing {}
    }
    new FS.File {
        path = "/etc/apt/trusted.gpg.d/zerotier.com.gpg"
        state = new FS.Missing {}
    }
}

function enable(): Listing<Resources.Resource> = new Listing {
    ...install()
    // todo
}
```

And now add it to our resource file:

```pkl
import "../modules/zerotier.pkl"

resources = new Listing {
    ...Apt.install(new Listing {
        // tools
        "tmux" "vim""git"  "mc" "nnn" "ncdu"
        "btop" "htop" "iftop"
        "kitty-terminfo"

        // services
        "ssh" "syncthing"
    })
    ...zerotier.enable()
}
```

Let's check the state:

```
Target: &{hamd:22 mike /home/mike/.ssh/id_ed25519}
Checking...
Resource Id                                  | Current State     | Expected State                                     
-----------------------------------------------------------------------------------------
apt:tmux                                     | 3.3a-3            | installed
apt:vim                                      | 2:9.0.1378-2      | installed
apt:git                                      | 1:2.39.2-1.1      | installed
apt:mc                                       | 3:4.8.29-2        | installed
apt:nnn                                      | 4.7-1             | installed
apt:ncdu                                     | 1.18-0.2          | installed
apt:btop                                     | 1.2.13-1          | installed
apt:htop                                     | 3.2.2-2           | installed
apt:iftop                                    | 1.0~pre4-9        | installed
apt:kitty-terminfo                           | 0.26.5-5          | installed
apt:ssh                                      | 1:9.2p1-2+deb12u2 | installed
apt:syncthing                                | 1.19.2~ds1-1+b4   | installed
file:/etc/apt/trusted.gpg.d/zerotier.com.gpg | missing           | 4b9b800c root:root 644
file:/etc/apt/sources.list.d/zerotier.list   | missing           | b71efca1 root:root 644
apt:zerotier-one                             | missing           | installed
```

Let's introduce a new command `actions` that
displays only what needs to be done:

```
+file:/etc/apt/trusted.gpg.d/zerotier.com.gpg                                                                         
+file:/etc/apt/sources.list.d/zerotier.list
+apt:zerotier-one
```

And apply:

```
Checking...
Applying...                                                                                                           
✓ +file:/etc/apt/trusted.gpg.d/zerotier.com.gpg                                                                       
✓ +file:/etc/apt/sources.list.d/zerotier.list                                                                         
✓ +apt:zerotier-one                                                                                                   
```

Now rerun `actions` to see that the system is fully
up-to-date:

```
Target:  declpi:22
Checking...
No actions necessary                                                                                                        
```
